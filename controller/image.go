package controller

import (
	"context"
	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

// DownloadDockerImage downloads a Docker image
func DownloadDockerImage(name string) (err error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DockerApiUrl))
	if err != nil {
		return err
	}
	defer func(cli *client.Client) {
		errClose := cli.Close()
		if errClose != nil {
			err = errClose
		}
	}(cli)

	reader, err := cli.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer func(reader io.ReadCloser) {
		errClose := reader.Close()
		if errClose != nil {
			err = errClose
		}
	}(reader)

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}

	return nil
}
