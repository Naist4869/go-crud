package service

import (
	"fmt"
	"go-crud/cache"
	"go-crud/model"
	"go-crud/serializer"

	"strings"
)

// DailyRankService 每日排行的服务
type DailyRankService struct {
}

// Get 获取排行
func (service *DailyRankService) Get() serializer.Response {
	var videos []model.Video

	// 从redis读取点击前十的视频
	vids, _ := cache.RedisClient.ZRange(cache.DailyRankKey, 0, 9).Result()

	if len(vids) > 0 {
		order := fmt.Sprintf("FIELD(id, %s) DESC", strings.Join(vids, ","))
		err := model.DB.Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	} else {
		err := model.DB.Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
