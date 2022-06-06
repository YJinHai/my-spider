package main

import (
	"github.com/gin-gonic/gin"
	"my-spider/internal/app/blog"
	"my-spider/pkg/config/env"
)

func main() {
	gin.SetMode(env.Conf.RunMode)

	gin.DisableConsoleColor()

	app := gin.New()
	//app.Use(middleware.Logger())
	//app.Use(middleware.Recovery())

	app.Use(gin.Logger(), gin.Recovery())
	blog.Router(app)

	app.Run("127.0.0.1:" + env.Conf.Port)

	//endless.ListenAndServe(":"+env.Conf.Port, app)
	/*	if err != nil {
		panic(err)
	}*/
}
