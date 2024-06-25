package routes

import (
	v1 "backend/api/v1"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

// InitRoute 初始化路由
func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
