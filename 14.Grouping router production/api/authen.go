package authen

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func Register(c *gin.Context) {
	c.String(http.StatusOK, "register")
}
