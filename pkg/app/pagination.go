package app

import (
	"blog_server/global"
	"blog_server/pkg/convert.go"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	pageSize, _ := convert.StrTo(c.Query("page_size")).Int()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
