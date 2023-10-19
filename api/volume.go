package api

import (
	"github.com/Juli-EXP/docker-volume-backup/backup"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func VolumeRouter(g *gin.RouterGroup) {
	g.GET("/", func(ctx *gin.Context) {
		code, res, err := getDockerVolumesWithSize()
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(code, res)
	})

	g.GET("/:name", func(ctx *gin.Context) {
		code, res, err := getDockerVolumeWithSize(ctx.Param("name"))
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(code, res)
	})
}

func getDockerVolumesWithSize() (code int, volumesResponse backup.VolumesResponse, err error) {
	volumesResponse, err = backup.GetDockerVolumesWithSize()
	if err != nil {
		return 500, backup.VolumesResponse{}, err
	}

	// Check if volumesResponse is empty
	if len(volumesResponse.Volumes) == 0 {
		return 204, volumesResponse, nil
	}

	return 200, volumesResponse, nil
}

func getDockerVolumeWithSize(volumeName string) (code int, response any, err error) {
	volumesResponse, err := backup.GetDockerVolumesWithSize()
	if err != nil {
		return 500, backup.Volume{}, err
	}

	// Check if volumesResponse is empty
	if len(volumesResponse.Volumes) == 0 {
		return 204, backup.Volume{}, nil
	}

	for _, volume := range volumesResponse.Volumes {
		if volume.Name == volumeName {
			return 200, volume, nil
		}
	}

	return 404, backup.Volume{}, errors.Errorf("No volume with the name %s was found", volumeName)
}
