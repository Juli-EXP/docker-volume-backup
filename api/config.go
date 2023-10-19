package api

import (
	"github.com/gin-gonic/gin"
)

func ConfigRouter(g *gin.RouterGroup) {
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "config",
		})
	})
}
