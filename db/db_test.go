package db

import (
	"log"
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	ResetDB()
	defer SQLDB.Close()
	defer os.Remove("./db.sqlite")

	//Add a Query
	log.Println("Add A Query")
	q := Query{}
	q.Text = "test"
	q.Type = "search"
	key, err := AddQuery(q)
	if err != nil {
		log.Println(err)
	}
	log.Println(key)

	//get queue
	log.Println("Print Queue")
	queue, err := GetQueue(10)
	if err != nil {
		log.Println(err)
	}
	log.Println(queue)

	//check resolved
	log.Println("Check if Resolved (shouldnt be)")
	q, err = CheckQuery(key)
	if err != nil {
		log.Println(err)
	}
	log.Println(q)

	//answer query
	log.Println("Answer the query")
	err = AnswerQuery(key, "42", "[\"facebook.com\"]")
	if err != nil {
		log.Println(err)
	}

	//get queue
	log.Println("Print Queue")
	queue, err = GetQueue(10)
	if err != nil {
		log.Println(err)
	}

	//check resolved
	log.Println("Check if Resolved (should be)")
	q, err = CheckQuery(key)
	if err != nil {
		log.Println(err)
	}
	log.Println(q)
	//recently resolved
	log.Println("Recently Resolved")
	resolved, err := GetRecent(10)
	if err != nil {
		log.Println(err)
	}
	log.Println(resolved)

}
