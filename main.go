package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello gin.")
	})
	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})

	app.StaticFile("favicon.ico", "./favicon.ico")
	app.Run(":80")

}

// func hello(c *gin.Context) {
// 	c.String(http.StatusOK, "hello,gin!")
// }
