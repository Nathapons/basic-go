package api

import "github.com/gin-gonic/gin"

func DownloadFile(c *gin.Context) {
	c.Header("Content-Description", "Simulation File Download")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+"download.jpg")
	c.Header("Content-type", "application/octet-stream")

	c.File("static/test.jpg")
}
