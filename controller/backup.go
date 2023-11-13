package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type CreateBackupOptions struct {
	VolumeName       string // Volume to be backed up
	BackupVolumeName string // Volume where the backup will be saved
	// TODO move to config.variables.go
	IncludeNfs  bool // Default false
	IncludeCifs bool // Default false
}

type DeleteBackupOptions struct{}

// CreateDockerVolumeBackup creates a backup of a Docker volume
func CreateDockerVolumeBackup(options CreateBackupOptions) (err error) {
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

	err = DownloadDockerImage(config.BackupContainerImage)
	if err != nil {
		return err
	}

	containerConfig := &container.Config{
		Image: config.BackupContainerImage,
		Labels: map[string]string{
			"dvb.container.temp": "true",
		},
		Cmd: []string{
			"ash",
			"-c",
			fmt.Sprintf("tar czf /dest/%s-%s.backup.tar.gz /data/.", options.VolumeName, fmt.Sprint(time.Now().Unix())),
		},
	}

	hostConfig := &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/data", options.VolumeName),
			fmt.Sprintf("%s:/dest", options.BackupVolumeName),
		},
	}

	// Create the container
	resp, err := cli.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		nil,
		nil,
		fmt.Sprintf("dvb-backup-%s", fmt.Sprint(time.Now().Unix())))
	if err != nil {
		return err
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	// Remove the container (cleanup)
	err = cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

// DeleteDockerVolumeBackup deletes a backup of a Docker volume
func DeleteDockerVolumeBackup(options DeleteBackupOptions) (err error) {
	//TODO
	return nil
}
