package api

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.POST("/add", AddTodo())
	r.POST("/update", UpdateTodo())
	r.GET("/view", ViewTodos())
	r.POST("/delete", DeleteTodo())
}
