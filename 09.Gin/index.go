package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Return plan text
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("Root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("profile"))
	})

	// Query params in URL
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.Run()
}
