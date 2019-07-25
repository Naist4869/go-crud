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

// DeleteVideo 删除视频接口
func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}

// VideoList 视频列表接口
func VideoList(c *gin.Context) {
	var service service.VideoListService
	res := service.List()
	c.JSON(200, res)

}

// ShowVideo 获取视频详情接口
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	res := service.Show(c.Param("id"))
	c.JSON(200, res)

}

// UpdateVideo 更新视频视频接口
func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
