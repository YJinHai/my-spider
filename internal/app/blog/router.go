package blog

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	apiRouter := router.Group("/api/blog")
	{
		apiRouter.GET("/list", GetAllBlog)
	}
}
