package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
	//driver for sqlite
	_ "github.com/mattn/go-sqlite3"
)

//SQLDB database
var SQLDB *sql.DB

//SQLPath path do db
var SQLPath string

//Query type
type Query struct {
	Key      int64     `json:"key"`
	Text     string    `json:"text"`
	Type     string    `json:"type"`
	Answer   string    `json:"answer"`
	List     []string  `json:"list"`
	ListJ    string    `json:"-"`
	Position int64     `json:"-"`
	Token    string    `json:"token"`
	Priority time.Time `json:"-"`
	Paid     bool      `json:"paid"`
}

//Charge type
type Charge struct {
	ID    string `json:"charge"`
	Query int64  `json:"query"`
}

//Init init sql
func Init() {
	var err error
	var created = false
	//check for the sqlite file
	if _, err := os.Stat("./db.sqlite"); os.IsNotExist(err) {
		log.Println("Creating database file")
		_, err := os.Create("./db.sqlite")
		if err != nil {
			log.Fatal("Failed to create SQLite file")
		}
		created = true
	}
	//Prepare DB
	SQLPath = "./db.sqlite"
	SQLDB, err = sql.Open("sqlite3", SQLPath)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Database setup failed.")
	}
	//create tables if necessary
	if created {
		CreateTables()
	}
	//load links
	initLinks()
}

//ResetDB remove existing database and recreate tables
func ResetDB() {
	//remove existing dataBasePath file
	os.Remove("./db.sqlite")
	Init()
}

//CreateTables create the user and posts tables
func CreateTables() {
	createTables := `
	CREATE TABLE queries (
		key integer primary key autoincrement,
		text varchar(255) not null,
		type varchar(20) not null,
		priority timestamp
	);
	CREATE TABLE resolved (
		key integer primary key not null,
		text varchar(255) not null,
		type varchar(20) not null,
		answer varchar(800) null,
		list varchar(2400) null
	);
	CREATE TABLE charges (
		key varchar(255) primary key
	);
	DELETE from queries;
  DELETE from resolved;
	DELETE from charges;
	`
	//Create the users table
	_, err := SQLDB.Exec(createTables)
	if err != nil {
		log.Println(err)
		log.Fatal("Unable to create database")
	}
	log.Println("Created SQL Tables")
}

//AddQuery add a new query to the db
func AddQuery(q Query) (int64, error) {
	//prepare the insert statement for adding a user
	insert, err := SQLDB.Prepare("INSERT into queries(text, type, priority) values(?, ?, ?)")
	if err != nil {
		return 0, errors.New("Error creating Query insert")
	}
	//insert
	result, err := insert.Exec(q.Text, q.Type, nil)
	if err != nil {
		return 0, errors.New("Failed to add Query")
	}
	//get ID to return
	id, err := result.LastInsertId()
	return id, nil
}

//GetQueue get the n top queries in the queue
func GetQueue(num int) ([]Query, error) {
	queries := []Query{}
	q := Query{}
	var priority sql.NullString
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type,priority FROM queries ORDER BY datetime(priority) DESC LIMIT (?)", num)
	defer rows.Close() //close query connection when function returns
	if err != nil {
		return queries, errors.New("Error getting queue")
	}
	for rows.Next() {
		q = Query{}
		err = rows.Scan(&q.Key, &q.Text, &q.Type, &priority)
		if err != nil {
			return queries, errors.New("Error scanning row")
		}
		// set the priority timestamp if the user has paid
		if priority.Valid {
			q.Paid = true
		}
		queries = append(queries, q)
	}

	return queries, nil
}

//IsDuplicate check if a text string has been answered
func IsDuplicate(text string) (bool, Query) {
	q := Query{}
	//select from resolved
	err := SQLDB.QueryRow("SELECT * FROM resolved WHERE text=(?)", text).Scan(&q.Key, &q.Text, &q.Type, &q.Answer, &q.ListJ)
	if err == nil {
		return true, q
	}
	return false, q
}

//FastForward auto answer questions that are old
func FastForward() {
	var queries []Query
	count := 0
	log.Println("Fast Forwarding")
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type FROM queries")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		q := Query{}
		err := rows.Scan(&q.Key, &q.Text, &q.Type)
		if err == nil {
			queries = append(queries, q)
		}
	}
	rows.Close()
	for i := 0; i < len(queries)-10; i++ {
		err := AnswerQuery(queries[i].Key, "Hey this is Jimmy, I'm a few thousand searches behind and pasting this answer in. Ask again if you want this answered.", "[\"http://gizoogle.com\", \"http://isitchristmas.com\"]")
		if err != nil {
			log.Println(err)
		}
		count++
	}
	log.Println(strconv.Itoa(count) + " queries auto answered")
}

//AnswerDuplicates looks for duplicate queries
func AnswerDuplicates() {
	count := 0
	log.Println("Removing duplicates")
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type FROM queries")
	var answers []Query
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		q := Query{}
		err := rows.Scan(&q.Key, &q.Text, &q.Type)
		if err == nil {
			isDupe, d := IsDuplicate(q.Text)
			if isDupe {
				a := Query{}
				a.Key = q.Key
				a.Answer = d.Answer
				a.ListJ = d.ListJ
				answers = append(answers, a)
				count++
			}
		}
	}
	rows.Close()
	for i := 0; i < len(answers); i++ {
		err := AnswerQuery(answers[i].Key, answers[i].Answer, answers[i].ListJ)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(strconv.Itoa(count) + " duplicates removed")
}

//AnswerQuery move a query to the resolved table with jimmy's answer
func AnswerQuery(key int64, answer string, list string) error {
	q := Query{}

	tx, err := SQLDB.Begin() //start transaction
	if err != nil {
		return errors.New("Could not establish transaction")
	}
	//get the query
	err = tx.QueryRow("SELECT text,type FROM queries WHERE key=(?)", key).Scan(&q.Text, &q.Type)
	if err != nil {
		tx.Rollback()
		return errors.New("Could not find query")
	}
	//add to resolved
	insert, err := tx.Prepare("INSERT into resolved(key, text, type, answer, list) values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return errors.New("Error creating Resolved insert")
	}
	_, err = insert.Exec(key, q.Text, q.Type, answer, list)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to add to resolved")
	}
	//delete from queries
	delete, err := tx.Prepare("DELETE from queries WHERE key=(?)")
	_, err = delete.Exec(key)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to delete query")
	}
	tx.Commit()
	return nil
}

//AddCharge adds a charge to the database and errors if charge exists
func AddCharge(key string) error {
	var c Charge
	err := SQLDB.QueryRow("SELECT key FROM charges WHERE key=?", key, 1).Scan(&c.ID)

	if err == nil {
		return errors.New("Charge already exists")
	}

	insert, err := SQLDB.Prepare("INSERT into charges(key) values(?)")
	if err != nil {
		return errors.New("Error creating Charge insert")
	}
	//insert
	_, err = insert.Exec(key)
	if err != nil {
		return errors.New("Failed to add Charge")
	}

	return nil
}

//MoveToFront moves a query to the front of the queue
func MoveToFront(key int64) error {
	update, err := SQLDB.Prepare("UPDATE queries SET priority=? WHERE key=?")
	if err != nil {
		return err
	}

	timestamp := time.Now()
	_, err = update.Exec(timestamp, key)
	log.Println("A user has paid.")

	if err != nil {
		return errors.New("Failed to move query to the pront of the queue")
	}

	return nil
}

//getQueuePosition return the queue position of a query
func getQueuePosition(key int64) (int64, error) {
	var err error
	var count int64
	var priority sql.NullString

	SQLDB.QueryRow("SELECT priority FROM queries WHERE key = (?)", key).Scan(&priority)

	if priority.Valid {
		// This is a paid query and can be compared as such
		err = SQLDB.QueryRow("SELECT COUNT(1) FROM queries WHERE datetime(priority) > datetime((?))", priority).Scan(&count)
	} else {
		// This is an unpaid query and need to be
		err = SQLDB.QueryRow("SELECT COUNT(1) FROM queries WHERE priority IS NOT NULL OR key < (?)", key).Scan(&count)
	}

	if err != nil {
		log.Println(err)
	}

	return count, nil
}

//CheckQuery see if a query is resolved and return the answer
func CheckQuery(key int64) (Query, error) {
	var list string
	q := Query{}
	//select from resolved
	err := SQLDB.QueryRow("SELECT * FROM resolved WHERE key=(?)", key).Scan(&q.Key, &q.Text, &q.Type, &q.Answer, &list)
	if err != nil {
		//get position
		q.Position, err = getQueuePosition(key)
		if err != nil {
			//we did not get the position
			return q, errors.New("Error getting queue position")
		}
		return q, errors.New("Query is not resolved")
	}
	//convert list to actual list
	err = json.Unmarshal([]byte(list), &q.List)
	if err != nil {
		q.List = nil
	}
	return q, nil
}

//GetRecent get the recently resolved posts
func GetRecent(num int) ([]Query, error) {
	var list string
	resolved := []Query{}
	r := Query{}
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type,answer,list FROM resolved ORDER BY key DESC LIMIT (?)", num)
	defer rows.Close() //close query connection when function returns
	if err != nil {
		return resolved, errors.New("Error getting recently resolved")
	}
	for rows.Next() {
		r = Query{}
		err = rows.Scan(&r.Key, &r.Text, &r.Type, &r.Answer, &list)
		if err != nil {
			return resolved, errors.New("Error scanning row")
		}
		//convert list to actual list
		err = json.Unmarshal([]byte(list), &r.List)
		if err != nil {
			r.List = nil
		}
		resolved = append(resolved, r)
	}

	return resolved, nil
}

//GetQuestion get a question by query id
func GetQuestion(key int64) (Query, error) {
	q := Query{}
	//get the query
	err := SQLDB.QueryRow("SELECT key,text,type FROM queries WHERE key=(?)", key).Scan(&q.Key, &q.Text, &q.Type)
	if err != nil {
		//not found in queries table so check resolved table
		err := SQLDB.QueryRow("SELECT key,text,type FROM resolved WHERE key=(?)", key).Scan(&q.Key, &q.Text, &q.Type)
		if err != nil {
			//the query is not in either table
			return q, errors.New("Could not find question")
		}
	}
	return q, nil
}
