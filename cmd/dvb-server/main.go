package main

import (
	"github.com/Juli-EXP/docker-volume-backup/api"
	"github.com/gin-gonic/gin"
)

func main(){
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	apiRouter := server.Group("/api")

	api.ConfigRouter(apiRouter.Group("/config"))
	api.VolumeRouter(apiRouter.Group("/volume"))
	api.BackupRouter(apiRouter.Group("/backup"))

	server.Run(":3000")
}