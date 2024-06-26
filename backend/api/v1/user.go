package v1

import (
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User

	_ = c.ShouldBindJSON(&data)
	//log.Println(data.Username)
	//log.Println(data.Password)
	//log.Println(data.Role)

	code = model.CheckUser(data.Username)

	//成功才写入数据库
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	//字符串转为数字
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//gorm规定，-1表示不作限制，查询所有
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{"status": code,
		"data":    data,
		"message": errmsg.GetErrMsg(code)})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
