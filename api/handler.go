package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTodo Todo Add Handler
func AddTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Todo{}
		c.Bind(&requestBody)
		err := SaveTodoToDB(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "not saved")
		} else {
			c.JSON(http.StatusOK, "added successfully")
		}
	}
}

// UpdateTodo Todo Update Handler
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Todo{}
		c.Bind(&requestBody)
		err := UpdateTodoFromDB(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "not updated")
		} else {
			c.JSON(http.StatusOK, "updated successfully")
		}
	}
}

// DeleteTodo Todo Delete Handler
func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Todo{}
		c.Bind(&requestBody)
		err := DeleteTodoFromDB(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "not deleted")
		} else {
			c.JSON(http.StatusOK, "deleted successfully")
		}
	}
}

// ViewTodos Todo View Handler
func ViewTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := GetAllTodosFromDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "not deleted")
		} else {
			c.JSON(http.StatusOK, data)
		}
	}
}
