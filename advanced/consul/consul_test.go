package consul

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestConsulName(t *testing.T) {
	r := gin.Default()

	// consul健康检查回调函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	go http.ListenAndServe(":8088", r)
	// 注册服务到consul
	RegistryConsul()

	// 从consul中发现服务
	consulFindServer()

	consulCheckHeath()
	consulKVTest()
	// 取消consul注册的服务
	// ConsulDeRegister()
	var str string
	fmt.Scan(&str)
}
