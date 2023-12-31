package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	count := 0
	runningDir, _ := os.Getwd()
	errorfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errorfile
	gin.DefaultWriter = accesslogfile
	r.Use(gin.Logger())

	// Return plan text
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("Root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		count = count + 1
		accesslogfile.WriteString(fmt.Sprintf("Count %d \n", count))
		c.Data(http.StatusOK, "text/html: charset=utf-8", []byte("profile"))
	})

	r.Run()
}
