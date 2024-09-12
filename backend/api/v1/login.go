package v1

import (
	"backend/middleware"
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var code int
	var token string
	_ = c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

func UserLogin(c *gin.Context) {
	var data model.User
	var code int
	var token string
	_ = c.ShouldBindJSON(&data)
	code = model.UserCheckLogin(data.Email, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
