package main

import (
	"eden/api/router"
	"eden/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	//db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func main() {
	configPath := "data/config.yml"
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()

	_ = web.Run(":" + conf.Server.Port)
}
