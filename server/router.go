package server

import (
	"go-crud/api"
	"go-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//创建视频
		v1.POST("video", api.CreateVideo)
		//删除视频
		v1.DELETE("video/:id", api.DeleteVideo)
		//获取视频列表
		v1.GET("videos", api.VideoList)
		//获取视频详情
		v1.GET("video/:id", api.ShowVideo)
		//更新视频
		v1.PATCH("video/:id", api.UpdateVideo)
		//请求Token
		v1.POST("upload/token", api.UploadToken)

		authed := r.Group("/")

		// 需要登录保护的
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
