package api

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

// to establish conn with db
func InitDB() {
	session, err := mgo.Dial("mongodb://myuser:secret@0.0.0.0:27017/todo_db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	} else {
		fmt.Println("successfully connected to database")
	}

	db = session.DB("")
}

// getting object with target collection
func DB() *mgo.Collection {
	return db.C("todos")
}
