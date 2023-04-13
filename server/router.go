package server

import (
	// "os"
	"github.com/gin-gonic/gin"
	// "net/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_api_bootstrap/docs"
	"gin_api_bootstrap/api"
	// "gin_api_bootstrap/middleware"
)

func NewRouter() *gin.Engine {
    router := gin.Default()

	// 使用 Logger 中间件，记录每个请求的日志信息
	router.Use(gin.Logger())

	// 简单的路由组: /api/v1
	api_v1 := router.Group("/api/v1")
	{
		api_v1.GET("/ping", api.Ping)
	}

    docs.SwaggerInfo.Title = "接口文档"
    docs.SwaggerInfo.Description = "接口文档"
    docs.SwaggerInfo.Version = "1.0"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
