package v1

import (
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("id"))
	email := c.Param("email")
	data, code := model.GetProfile(email)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code)})
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	//id, _ := strconv.Atoi(c.Param("id"))
	email := c.Param("emails")

	_ = c.ShouldBindJSON(&data)
	//log.Println(data)
	code = model.UpdateProfile(email, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
