package main

import (
	"gin_api_bootstrap/server"
	"gin_api_bootstrap/config"
	"gin_api_bootstrap/util"
	"gin_api_bootstrap/query"
)

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main() {
	config.Init()
	// 连接数据库
	query.Database(config.Config.DBURL)
	// 创建路由服务
	router := server.NewRouter()
	url := config.Config.Host + ":" + util.Int2Str(config.Config.Port)
	router.Run(url)
}
