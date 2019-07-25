package api

import (
	"go-crud/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 创建视频接口
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
