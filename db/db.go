package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	//driver for sqlite
	_ "github.com/mattn/go-sqlite3"
)

//SQLDB database
var SQLDB *sql.DB

//SQLPath path do db
var SQLPath string

//Query type
type Query struct {
	Key    int64  `json:"key"`
	Text   string `json:"text"`
	Type   string `json:"type"`
	Answer string `json:"answer"`
}

//InitDB init sql
func InitDB() {
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
}

//ResetDB remove existing database and recreate tables
func ResetDB() {
	//remove existing dataBasePath file
	os.Remove("./db.sqlite")
	InitDB()
}

//CreateTables create the user and posts tables
func CreateTables() {
	createTables := `
	CREATE TABLE queries (
		key integer primary key autoincrement,
		text varchar(255) not null,
		type varchar(60) not null
	);
	CREATE TABLE resolved (
		key integer primary key not null,
		text varchar(255) not null,
		type varchar(60) not null,
		answer varchar(500) null
	);
	DELETE from queries;
    DELETE from resolved;
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
	insert, err := SQLDB.Prepare("INSERT into queries(text, type) values(?, ?)")
	if err != nil {
		return 0, errors.New("Error creating Query insert")
	}
	//insert
	result, err := insert.Exec(q.Text, q.Type)
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
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type FROM queries LIMIT (?)", num)
	defer rows.Close() //close query connection when function returns
	if err != nil {
		return queries, errors.New("Error getting queue")
	}
	for rows.Next() {
		err = rows.Scan(&q.Key, &q.Text, &q.Type)
		if err != nil {
			return queries, errors.New("Error scanning row")
		}
		queries = append(queries, q)
	}

	return queries, nil
}

//AnswerQuery move a query to the resolved table with jimmy's answer
func AnswerQuery(key int64, answer string) error {
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
	insert, err := tx.Prepare("INSERT into resolved(key, text, type, answer) values(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return errors.New("Error creating Resolved insert")
	}
	_, err = insert.Exec(key, q.Text, q.Type, answer)
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

//CheckQuery see if a query is resolved and return the answer
func CheckQuery(key int64) (Query, error) {
	q := Query{}

	//select from resolved
	err := SQLDB.QueryRow("SELECT * FROM resolved WHERE key=(?)", key).Scan(&q.Key, &q.Text, &q.Text, &q.Answer)
	if err != nil {
		return q, errors.New("Query is not resolved")
	}
	return q, nil
}

//GetRecent get the recently resolved posts
func GetRecent(num int) ([]Query, error) {
	resolved := []Query{}
	r := Query{}
	//create sql query
	rows, err := SQLDB.Query("SELECT key,text,type,answer FROM resolved ORDER BY key DESC LIMIT (?)", num)
	defer rows.Close() //close query connection when function returns
	if err != nil {
		return resolved, errors.New("Error getting recently resolved")
	}
	for rows.Next() {
		err = rows.Scan(&r.Key, &r.Text, &r.Type, &r.Answer)
		if err != nil {
			return resolved, errors.New("Error scanning row")
		}
		resolved = append(resolved, r)
	}

	return resolved, nil
}
