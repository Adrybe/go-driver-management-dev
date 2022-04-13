package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	port = ":" + port

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "message : pong")
	})

	r.Run()
}
