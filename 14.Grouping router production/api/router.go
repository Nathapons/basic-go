package api

import "github.com/gin-gonic/gin"

func SetUp(router *gin.Engine) {
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("list", ListProduct)
		stockAPI.GET("create", CreateProduct)
	}

	fileAPI := router.Group("/file")
	{
		fileAPI.GET("download", DownloadFile)
	}
}
