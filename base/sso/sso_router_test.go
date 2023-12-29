package sso

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestSSORouter(t *testing.T) {
	r := gin.Default()

	r.POST("/sso", beforeLogin)
	r.POST("/api/login", login)
	r.Run(":8888")
}
