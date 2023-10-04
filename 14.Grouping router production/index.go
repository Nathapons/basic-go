package main

import (

	"github.com/gin-gonic/gin"
	"demo14/api"
)

func main() {
	router := gin.Default()
	api.SetUp(router)
	router.Run()
}
