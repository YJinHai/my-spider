package blog

import (
	"github.com/gin-gonic/gin"
	"my-spider/internal/pkg/blog"
	"my-spider/pkg/response"
)

func Jiajun(c *gin.Context) {
	obj := blog.NewJiajunInfo()
	obj.SetWebPage()
	retData := obj.GetData()
	list := []interface{}{retData, retData, retData, retData, retData, retData, retData, retData}
	response.ApiReturn(1, list, c)
}

func JiajunPage() interface{} {
	obj := blog.NewJiajunInfo()
	obj.SetWebPage()
	retData := obj.GetData()
	return retData
}
