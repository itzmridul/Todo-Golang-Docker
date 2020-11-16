package main

import (
	"HelloWorld/api"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Gin-Gonic Server")

	api.InitDB()

	startServer()
}

func startServer() {
	// creating an instance of gin
	router := gin.Default()

	// dummy url for check if server is running
	router.GET("/", checkStatus())

	//creating the routes
	api.Init(router)

	// creating http server instance
	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}

	//starting server
	s.ListenAndServe()
}

// func for checking server status
func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
