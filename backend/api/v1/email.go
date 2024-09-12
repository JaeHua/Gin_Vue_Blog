package v1

import (
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// GetValidateCode 获取验证码
func GetValidateCode(c *gin.Context) {
	rd := model.GetRedis()
	// 获取目的邮箱
	//获取前端发来的验证码
	var mail = model.MailCode{}
	_ = c.ShouldBind(&mail)
	log.Println(mail.Mail)
	em := []string{mail.Mail}
	vCode, code, err := model.SendEmailValidate(em)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 验证码存入redis 并设置过期时间5分钟,注意key的唯一性
	parts := strings.Split(em[0], "@")
	err1 := rd.Db.Set("vCode"+parts[0], vCode, 300000000000).Err()

	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  4002,
			"message": errmsg.GetErrMsg(4002),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": "验证码发送成功",
		"vcode":   vCode,
	})
}

// ValidateEmailCode 验证邮箱
func ValidateEmailCode(c *gin.Context) {
	rd := model.GetRedis()
	//获取前端发来的验证码
	var em = model.MailCode{}
	_ = c.ShouldBind(&em)
	vCode := em.VCode
	// 获取存储在redis中的验证码
	key := strings.Split("vCode"+em.Mail, "@")
	//log.Println(key[0])
	rdVCode := rd.Db.Get(key[0]).Val()
	//log.Println("vcode" + vCode)
	//log.Println("edvcode" + rdVCode)
	if rdVCode != "" && vCode == rdVCode {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "验证成功",
		})
		return
	} else {
		if rdVCode != "" && vCode != rdVCode {
			code = 4003
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			return
		} else {
			code = 4004
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			return
		}
	}
}
