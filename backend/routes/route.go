package routes

import (
	v1 "backend/api/v1"
	"backend/middleware"
	"backend/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"

	"github.com/gin-contrib/static"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

// InitRoute 初始化路由
func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//!!!
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(static.Serve("/", static.LocalFile("web/front/dist", true)))
	r.Use(static.Serve("/admin", static.LocalFile("web/admin/dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("web/front/dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})

	/*
		后台管理路由接口
	*/
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
		//更新信息
		auth.PUT("profile/:emails", v1.UpdateProfile)
		auth.GET("user/info", v1.Info)
		auth.PUT("editProfile", v1.UserInfoEdit)
	}
	/*
		前端展示页面接口
	*/
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
		public.POST("userlogin", v1.UserLogin)
		public.GET("profile/:email", v1.GetProfile)
		public.POST("register/getcode", v1.GetValidateCode)
		public.POST("register/verify", v1.ValidateEmailCode)
		public.POST("register", v1.UserRegister)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
