package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"

	"gin_api_bootstrap/util"
)

type T_Config struct {
	Host     string
	Port     int
	ApiKey	 string
	DBURL    string
}

var Config T_Config

func Init() {
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config.Host = os.Getenv("HOST")
	Config.Port = util.Str2Int(os.Getenv("PORT"))
	Config.ApiKey = os.Getenv("APIKEY")
	Config.DBURL = os.Getenv("DBURL")
}
