package v1

import (
	"backend/middleware"
	"backend/model"
	"backend/utils/errmsg"
	"backend/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var code int
	var token string
	_ = c.ShouldBindJSON(&data)
	data.Email = data.Username
	code = model.CheckLogin(data.Email, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Email)
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
		token, code = middleware.SetToken(data.Email)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

func UserRegister(c *gin.Context) {
	var data model.User
	var profile model.Profile
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)
	profile.Name = data.Username
	profile.Email = data.Email
	data.Role = 2
	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}
	code = model.CheckUser(data.Username, data.Email)
	//成功才写入数据库
	if code == errmsg.ERROR_USERNAME_USED {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
		model.CreateProfile(&profile)
	}
	token, _ := middleware.SetToken(data.Email)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
