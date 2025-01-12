package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"} 
	config.AllowOrigins = []string{"http://localhost:8080","http://www.i4k.tv"}
	config.AllowOrigins = []string{"http://localhost:8080","https://www.i4k.tv"}
	config.AllowCredentials = true
	return cors.New(config)
}
