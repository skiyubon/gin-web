package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	g.GET("/", hello)
	g.Run(":80")

}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello,gin!")
}
