package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleBookRequest(c *gin.Context) {
	from, to, by := c.Param("from"), c.Param("to"), c.Param("by")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to, "by": by})
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func loginRequest(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		if form.Username == "admin" && form.Password == "1234" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Can not authorized"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind"})
	}
}

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

	// Parameter in URL
	r.GET("book/:from/:to/:by", handleBookRequest)

	r.POST("/login", loginRequest)

	r.Run()
}
