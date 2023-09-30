package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func register(c *gin.Context) {
	c.String(http.StatusOK, "register")
}

func listProduct(c *gin.Context) {
	c.String(http.StatusOK, "listProduct")
}

func createProduct(c *gin.Context) {
	c.String(http.StatusOK, "createProduct")
}

func main() {
	router := gin.Default()
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", login)
		authenAPI.GET("/register", register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("list", listProduct)
		stockAPI.GET("create", createProduct)
	}
	router.Run()
}
