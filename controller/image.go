package controller

import (
	"context"
	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
)

// DownloadDockerImage downloads a Docker image
func DownloadDockerImage(imageName string) (err error) {
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

	reader, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer func(reader io.ReadCloser) {
		errClose := reader.Close()
		if errClose != nil {
			err = errClose
		}
	}(reader)

	_, err = io.Copy(io.Discard, reader)
	if err != nil {
		return err
	}

	return nil
}

func CheckDockerImageLocal(imageName string) (imageExists bool, err error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DockerApiUrl))
	if err != nil {
		return false, err
	}
	defer func(cli *client.Client) {
		errClose := cli.Close()
		if errClose != nil {
			err = errClose
		}
	}(cli)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageName {
				return true, nil
			}
		}
	}

	return false, nil
}
