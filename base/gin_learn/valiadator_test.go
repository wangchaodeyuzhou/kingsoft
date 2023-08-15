package gin_learn

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
	"time"
)

type Info struct {
	CreateTime time.Time `form:"create_time" binding:"required,timing" time_format:"2006-01-02"`
	UpdateTime time.Time `form:"update_time" binding:"required,timing" time_format:"2006-01-02"`
}

// 自定义验证规则断言
func timing(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func TestValidator(t *testing.T) {
	r := gin.Default()
	r.POST("/register", RegisterAccount)
	r.Run(":9999")
}

// 自定义验证
func TestField(t *testing.T) {
	r := gin.Default()
	// 注册验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", timing)
		if err != nil {
			println("success")
		}
	}

	r.GET("/time", getTime)
	r.Run(":9999")
}

func TestValidatorRegister(t *testing.T) {
	validateMy = validator.New()
	r := gin.Default()

	r.GET("/r", registerHandler)
	r.Run(":9999")
}

func getTime(c *gin.Context) {
	var b Info
	// 数据模型绑定查询字符串验证
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "param is error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "time are valid!"})
}
