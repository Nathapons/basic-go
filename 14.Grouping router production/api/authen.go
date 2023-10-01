package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var msg struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	msg.Username = "admin"
	msg.Password = "1234"
	c.JSON(http.StatusOK, msg)
}

func Register(c *gin.Context) {
	c.String(http.StatusOK, "register")
}
