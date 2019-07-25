package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// DeleteVideoService 创建视频的服务
type DeleteVideoService struct {
}

// Delete 删除视频函数
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50003,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Msg: "删除视频成功",
	}
}
