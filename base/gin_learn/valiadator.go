package gin_learn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validateMy *validator.Validate

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age" binding:"gte=1,lte=120"`
}

func RegisterAccount(c *gin.Context) {
	var req *RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}

	fmt.Println("register success")
	c.JSON(http.StatusOK, "success")
}

type RegisterData struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	Age      uint8  `json:"age" validate:"required,gte=3,lte=20"`
}

func registerHandler(c *gin.Context) {
	var data RegisterData
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("should bind json", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(data.Age, data.Password, data.Email, data.Username)

	// 使用数据验证库来验证数据
	if err := validateMy.Struct(data); err != nil {
		fmt.Println("err : ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

type RegisterPageData struct {
	Page     int `json:"age" validate:"gte=1,lte=10"`
	PageSize int `json:"pageSize" validate:"min=10,max=30"`
}

func registerPageData(c *gin.Context) {
	data := RegisterPageData{}
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("should bind json", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(data.Page, data.PageSize)
	if err := validateMy.Struct(data); err != nil {
		fmt.Println("err : ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
