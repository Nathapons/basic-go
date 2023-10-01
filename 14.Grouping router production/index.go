package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"demo14/api"
)

func main() {
	router := gin.Default()
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", api.Login)
		authenAPI.GET("/register", api.Register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("list", api.ListProduct)
		stockAPI.GET("create", api.CreateProduct)
	}
	router.Run()
}
