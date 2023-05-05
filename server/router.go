package server

import (
	"os"
	"io"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	// "net/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_api_bootstrap/docs"
	"gin_api_bootstrap/api"
	"gin_api_bootstrap/middleware"
)

func NewRouter() *gin.Engine {
    router := gin.Default()

	// 使用 Logger 中间件，记录每个请求的日志信息
	// router.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        //TODO 写入对应文件的逻辑
		log1, _ := os.Create("log.log")
		gin.DefaultWriter = io.MultiWriter(log1)
		// 输出自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))

	// 简单的路由组: /api/v1
	api_v1 := router.Group("/api/v1")
	{
		api_v1.GET("/ping", api.Ping)
		api_v1.GET("/checkApikey", middleware.AuthMiddleware(), api.CheckApikeyApi)
		api_v1.GET("/user/list", middleware.AuthMiddleware(), api.ListUserApi)
		api_v1.GET("/user/detail", middleware.AuthMiddleware(), api.GetUserDetailApi)
		api_v1.POST("/user/add", middleware.AuthMiddleware(), api.AddUserApi)
		api_v1.PUT("/user/update/:id", middleware.AuthMiddleware(), api.UpdateUserApi)
	}

    docs.SwaggerInfo.Title = "接口文档"
    docs.SwaggerInfo.Description = "接口文档"
    docs.SwaggerInfo.Version = "1.0"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
