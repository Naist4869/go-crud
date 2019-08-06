package api

import (
	"go-crud/service"

	"github.com/gin-gonic/gin"
)

// DailyRank 获取Token接口
func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
