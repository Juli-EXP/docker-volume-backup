package docker

import (
	"context"
	"encoding/json"

	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type VolumeType string

const (
	Nfs   VolumeType = "nfs"
	Cifs  VolumeType = "cifs"
	Local VolumeType = "local"
)

type VolumeInfo struct {
	Name string
	Type VolumeType
}

func GetDockerVolumes() ([]VolumeInfo, error) {
	var volumeInfoList []VolumeInfo

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DOCKER_API_URL))
	if err != nil {
		panic(err)
	}

	volumes, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Iterate over the list of volumes and get their options
	for _, volume := range volumes.Volumes {
		volumeInfo, err := cli.VolumeInspect(context.Background(), volume.Name)
		if err != nil {
			return nil, err
		}

		// Convert to json
		opt, err := json.Marshal(volumeInfo.Options)
		if err != nil {
			return nil, err
		}

		// Check if options type is nfs, cifs or empty
		var volumeType VolumeType
		if string(opt) != "null" {
			switch volumeInfo.Options["type"] {
			case "nfs":
				volumeType = Nfs
			case "cifs":
				volumeType = Cifs
			default:
				volumeType = Local
			}
		} else {
			volumeType = Local
		}

		// Append volume name and type to the result list
		volumeInfoList = append(volumeInfoList, VolumeInfo{
			Name: volume.Name,
			Type: volumeType,
		})
	}

	// Close the Docker client
	cli.Close()

	return volumeInfoList, nil
}
