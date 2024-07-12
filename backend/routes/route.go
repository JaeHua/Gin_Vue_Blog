package routes

import (
	v1 "backend/api/v1"
	"backend/middleware"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

// InitRoute 初始化路由
func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//!!!
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	//托管静态资源
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon.icon", "static/admin/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.PUT("userinfo/:id", v1.ResetUserInfo)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArt)

		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		//上传文件
		auth.POST("upload", v1.UpLoad)
	}
	public := r.Group("api/v1")
	{
		public.POST("user/add", v1.AddUser)
		public.GET("users", v1.GetUsers)
		public.GET("user/:id", v1.GetUserInfo)
		public.GET("category/:id", v1.GetCateInfo)
		public.GET("category", v1.GetCate)
		public.GET("article", v1.GetArt)
		public.GET("article/info/:id", v1.GetArtInfo)
		public.GET("category/article/:id", v1.GetCateArt)
		public.POST("login", v1.Login)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
