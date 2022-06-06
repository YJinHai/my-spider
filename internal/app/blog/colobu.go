package blog

import (
	"github.com/gin-gonic/gin"
	"my-spider/internal/pkg/blog"
	"my-spider/pkg/response"
)

func Colobu(c *gin.Context) {
	obj := blog.NewColobuInfo()
	obj.SetWebPage()
	retData := obj.GetData()
	list := []interface{}{retData}
	response.ApiReturn(1, list, c)
}

func ColobuPage() interface{} {
	obj := blog.NewColobuInfo()
	obj.SetWebPage()
	retData := obj.GetData()
	return retData
}
