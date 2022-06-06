package blog

import (
	"github.com/gin-gonic/gin"
	"my-spider/pkg/response"
)

func GetAllBlog(c *gin.Context) {
	var list []interface{}

	list = append(list, ColobuPage())
	list = append(list, JiajunPage())
	response.ApiReturn(1, list, c)
}
