package model

import ( "github.com/jinzhu/gorm"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)
// Video 模型
type Video struct {
	gorm.Model
	Title string
	Info  string
	Video string
	Poster string
}
	// PosterURL 获取视频封面地址
func (video *Video) PosterURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Poster,oss.HTTPGet,600)
	return signedGetURL


}
func (video *Video) VideoURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Video,oss.HTTPGet,600)
	return signedGetURL


}