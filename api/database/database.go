package database

import (
	// "fmt"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

var conn *dbr.Connection
var err error

// var eventLog EventLog

// var once sync.Once

func GetConnection() (*dbr.Session, error) {
	// conn, err = dbr.Open("sqlite3", "main.db", &eventLog)
	conn, err = dbr.Open("sqlite3", "database/main.db", nil)
	// why is defer auto-closing
	//	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sess := conn.NewSession(nil)
	// sess := conn.NewSession(&eventLog)

	return sess, err
}

// conn, _ := dbr.Open("sqlite3", "test.sqlite", nil)

// // create a session for each business unit of execution (e.g. a web request or goworkers job)
// sess := conn.NewSession(nil)

// get a record
// var suggestion Suggestion
// sess.Select("id", "title").From("suggestions").Where("id = ?", 1).Load(&suggestion)

// // JSON-ready, with dbr.Null* types serialized like you want
// json.Marshal(&suggestion)
