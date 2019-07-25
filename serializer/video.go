package serializer

import (
	"go-crud/model"
)
// Video 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}
// BuildVideo 序列化用户
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:  item.Title,
		Info:  item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}