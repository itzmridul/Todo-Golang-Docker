package api

import (
	uuid "github.com/satori/go.uuid"

	"gopkg.in/mgo.v2/bson"
)

// GetAllTodosFromDB - Mongo
func GetAllTodosFromDB() ([]Todo, error) {
	res := []Todo{}
	if err := DB().Find(nil).All(&res); err != nil {
		return []Todo{}, err
	}

	return res, nil
}

// SaveTodoToDB - Mongo
func SaveTodoToDB(todo Todo) error {
	uuidVar := uuid.NewV4() // generating unique id
	todo.Id = uuidVar.String()
	return DB().Insert(todo)
}

// DeleteTodoFromDB - Mongo
func DeleteTodoFromDB(todo Todo) error {
	return DB().Remove(bson.M{"_id": todo.Id})
}

// UpdateTodoFromDB - Mongo
func UpdateTodoFromDB(todo Todo) error {
	// update
	query := bson.M{
		"_id": todo.Id,
	}
	doc := bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"tags":        todo.Tags,
	}
	return DB().Update(query, doc)
}
