package backup

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
	IncludeNfs       bool
	IncludeCifs      bool
}

type DeleteBackupOptions struct{}

// Creates a backup of a Docker volume
func CreateDockerVolumeBackup(options CreateBackupOptions) error {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DOCKER_API_URL))
	if err != nil {
		return err
	}
	defer cli.Close()

	// TODO download image

	config := &container.Config{
		Image: config.BACKUP_CONTAINER_IMAGE,
		Labels: map[string]string{
			"com.dvb.container": "true",
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
		config,
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
			panic(err)
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

// Deletes a backup of a Docker volume
func DeleteDockerVolumeBackup(options DeleteBackupOptions) error {
	return nil
}
