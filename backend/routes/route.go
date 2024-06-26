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
		router.POST("category/add", v1.AddCate)
		router.GET("category", v1.GetCate)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
