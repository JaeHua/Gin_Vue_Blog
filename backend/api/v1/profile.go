package v1

import (
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code)})
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))

	_ = c.ShouldBindJSON(&data)
	log.Println(data)
	code = model.UpdateProfile(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
