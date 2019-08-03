package api

import (
	"go-crud/service"
"github.com/gin-gonic/gin"
)
	



// UploadToken 获取Token接口
func UploadToken(c *gin.Context) {
	 service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Token()
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}