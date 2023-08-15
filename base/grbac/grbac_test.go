package grbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
	"testing"
	"time"
)

func TestGrbac(t *testing.T) {
	r := gin.New()
	r.Use(Authentication())

	r.POST("/addAuth", AddAuth)

	r.Run(":8889")
}

func AddAuth(c *gin.Context) {
	fmt.Println("AddAuth")
}

func Authentication() gin.HandlerFunc {
	rbac, err := grbac.New(grbac.WithYAML("r.yaml", time.Minute*10))
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		// c.Request.Header.Set("Grant", "Root")
		state, err := rbac.IsRequestGranted(c.Request, []string{"Root"})
		if err != nil {
			c.Abort()
			fmt.Println("没有权限111")
			return
		}
		if state.IsGranted() {
			c.Next()
			fmt.Println("有权限2222")
			return
		}
		fmt.Println("没有权限444")
		c.Abort()
		return
	}
}
