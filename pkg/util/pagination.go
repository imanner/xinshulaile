package util

// 分页页码的获取方法 根据传入的参数和配置 获取当前页面的数量

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/imanner/gin-xinshulaila/pkg/setting"
)

// 分页页码的获取方法
func GetPage(c *gin.Context) int {
	// 设定默认值W为0
	result := 0
	// 获取page对应的值
	page, _ := com.StrTo(c.Query("page")).Int()
	// 如果有分页需求 获取页面的数量
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
