package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	runningDir, _ := os.Getwd()
	errorfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errorfile
	gin.DefaultWriter = accesslogfile
	r.Use(gin.Logger())

	// Return plan text
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("Root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("profile"))
	})

	r.Run()
}
