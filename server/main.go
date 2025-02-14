package main

import (
	"github.com/gin-gonic/gin" 
	// imports gin framework to create an HTTP server
)

func main() {
	//creates a new Gin router instance, like express.js for node.js
	r := gin.Default()

	//define a simple route for get requests at "/". c is a pointer to a gin.Context object
	r.GET("/", func(c *gin.Context) {
		//sends a json response with a status code of 200 ok
		c.JSON(200, gin.H{
		// gin.H is just a shorthand for map[string]interface{}.
		// It allows you to define JSON-like key-value pairs easily.
			"message" : "Hello world!",
		})
	})

	r.Run(":8080")

}