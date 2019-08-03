package service

import (
	"go-crud/serializer"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// "github.com/google/uuid"
)

// UploadTokenService 获得上传oss token的服务
type UploadTokenService struct {
	FileName string `form:"filename"  json:"filename"`
}

// Token 获取Token函数
func (service *UploadTokenService) Token() serializer.Response {

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
		oss.ContentType("image/jpeg"),
	}

	// key := "upload/poster/" + uuid.Must(uuid.NewRandom()).String() + ".jpeg"
	key := "upload/poster/" + service.FileName

	// 	签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600,options...)
	if err != nil {
		return serializer.Response{
			Status: 40002,
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


	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600,options...)
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}

}
