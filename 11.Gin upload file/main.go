package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	runningDir, _ := os.Getwd()
	var count int = 0
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		count = count + 1
		file, _ := c.FormFile("file")
		extension := filepath.Ext(file.Filename)

		c.SaveUploadedFile(file, fmt.Sprintf("%s/uploaded/%d%s", runningDir, count, extension))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run()
}
