package service

import (
	"go-crud/serializer"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// "github.com/google/uuid"
)

// UploadVideoService 获得上传oss token的服务
type UploadVideoService struct {
	FileName string `form:"filename"  json:"filename" binding:"required" `
}

// POST POST视频地址
func (service *UploadVideoService) POST() serializer.Response {

	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 40001,
			Msg:    "没取到Token",
			Error:  err.Error(),
		}
	}
	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "没取到Bucket",
			Error:  err.Error(),
		}
	}
		//带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType("video/mp4"),
		
	}

	// key := "upload/video/" + uuid.Must(uuid.NewRandom()).String() + ".jpeg"
	key := "upload/video/" + service.FileName

	// 	签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600,options...)
	if err != nil {
		return serializer.Response{
			Status: 40003,
			Msg:    "没取到签名",
			Error:  err.Error(),
		}
	}
	// err = bucket.PutObjectFromFileWithURL(signedPutURL,key)
	// if err != nil {
	// 	return serializer.Response{
	// 		Status: 40003,
	// 		Msg:    "没上传成功",
	// 		Error:  err.Error(),
	// 	}
	// }


	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			
		},
	}

}
