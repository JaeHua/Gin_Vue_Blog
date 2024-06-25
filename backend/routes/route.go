package routes

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRoute 初始化路由
func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hi", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
