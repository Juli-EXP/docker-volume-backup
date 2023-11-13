package main

import (
	"github.com/Juli-EXP/docker-volume-backup/api"
	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	apiRouter := server.Group("/api")

	api.ConfigRouter(apiRouter.Group("/config"))
	api.VolumeRouter(apiRouter.Group("/volume"))
	api.BackupRouter(apiRouter.Group("/controller"))

	err := server.Run(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
}
