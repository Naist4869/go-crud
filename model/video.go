package model

import (
	"go-crud/cache"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Video 模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	Video  string
	Poster string
}

// PosterURL 获取视频封面地址
func (video *Video) PosterURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Poster, oss.HTTPGet, 600)
	return signedGetURL
}

// VideoURL 获取视频地址
func (video *Video) VideoURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Video, oss.HTTPGet, 600)
	return signedGetURL

}

// View 获取点击数
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加点击数
func (video *Video) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	//增加排行榜点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
