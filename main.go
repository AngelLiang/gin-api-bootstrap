package main

import (
	"gin_api_bootstrap/server"
	"gin_api_bootstrap/config"
	"gin_api_bootstrap/util"
)

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main() {
	config.Init()
	router := server.NewRouter()
	url := config.Config.Host + ":" + util.Int2Str(config.Config.Port)
	router.Run(url)
}
