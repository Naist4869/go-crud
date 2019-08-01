package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Info string `form:"info" json:"info" binding:"required,max=80"`
	Poster	string `form:"poster" json:"poster"`

}

// Update 更新视频函数
func (service *UpdateVideoService) Update(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(&video).Update("info", service.Info).Error
	if err != nil {
		return serializer.Response{
			Status: 50004,
			Msg:    "视频更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Msg:  "更新视频成功",
	}
}
