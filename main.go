package main

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"Go-Online-Judge/models"
	"Go-Online-Judge/routers"
)

const (
	APP_CONFIG_FILE = "conf/app.yaml"
	LOG_CONFIG_FILE = "conf/log.xml"
)

func init() {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(APP_CONFIG_FILE)
	if e := viper.ReadInConfig(); e != nil {
		panic(e)
	}

	logger, err := log.LoggerFromConfigAsFile(LOG_CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)

	dbType := viper.GetString("database.type")
	dbURL := viper.GetString("database.url")
	models.InitDB(dbType, dbURL)

	if err := models.InitRedis(); err != nil {
		panic(err)
	}
}

func main() {
	router := routers.InitRouter()
	router.Run(":" + viper.GetString("app.port"))
}
