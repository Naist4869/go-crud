package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// CreateVideoService 创建视频的服务
type CreateVideoService struct {
	Title string `form:"Title" json:"Title" binding:"required,min=5,max=30"`
	Info  string `form:"Info" json:"Info" binding:"required,max=80"`
}

// Create 创建视频函数
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Msg:  "创建视频成功",
	}
}
