package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ShowVideoService 视频详情的服务
type ShowVideoService struct {
}

// Show 视频详情函数
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.Limit(10).First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系列问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
