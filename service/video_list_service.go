package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// VideoListService 视频列表的服务
type VideoListService struct {
}

// List 视频列表函数
func (service *VideoListService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
