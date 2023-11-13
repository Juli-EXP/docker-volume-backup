package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Volume struct {
	Name      string `json:"Name"`
	UsageData struct {
		Size int64 `json:"Size"`
	} `json:"UsageData"`
	Type   config.StorageType `json:"Type"`
	Labels map[string]string  `json:"Labels"`
	//TODO Backup status
}

type VolumesResponse struct {
	Volumes []Volume `json:"Volumes"`
}

// GetDockerVolumes returns array of all Docker volumes with Name and Type
func GetDockerVolumes() (volumeResponse VolumesResponse, err error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DockerApiUrl))
	if err != nil {
		return VolumesResponse{}, err
	}
	defer func(cli *client.Client) {
		errClose := cli.Close()
		if errClose != nil {
			err = errClose
		}
	}(cli)

	volumes, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return VolumesResponse{}, err
	}

	// Iterate over the list of volumes and get their options
	for _, dockerVolume := range volumes.Volumes {
		volumeInfo, err := cli.VolumeInspect(context.Background(), dockerVolume.Name)
		if err != nil {
			return VolumesResponse{}, err
		}

		// Convert to json
		opt, err := json.Marshal(volumeInfo.Options)
		if err != nil {
			return VolumesResponse{}, err
		}

		// Check if options type is nfs, cifs or empty
		var storageType config.StorageType
		if string(opt) != "null" {
			switch volumeInfo.Options["type"] {
			case "nfs":
				storageType = config.Nfs
			case "cifs":
				storageType = config.Cifs
			default:
				storageType = config.Local
			}
		} else {
			storageType = config.Local
		}

		// Append volume name and type to the result list
		volumeResponse.Volumes = append(volumeResponse.Volumes, Volume{
			Name: dockerVolume.Name,
			Type: storageType,
		})
	}

	return volumeResponse, nil
}

// GetDockerVolumesWithSize returns array of all Docker volumes with Name, Size, Type and Labels
func GetDockerVolumesWithSize() (VolumesResponse, error) {
	var url string
	var httpClient http.Client

	if strings.HasPrefix(config.DockerApiUrl, "unix://") {
		httpClient = http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", strings.TrimPrefix(config.DockerApiUrl, "unix://"))
				},
			},
		}
		url = "http://localhost/system/df"
	} else if strings.HasPrefix(config.DockerApiUrl, "http://") {
		httpClient = http.Client{}
		url = config.DockerApiUrl + "/system/df"
	} else {
		return VolumesResponse{}, errors.Errorf("Cannot determine the type of DockerApiUrl")
	}

	// Send an HTTP request to the Docker API
	resp, err := httpClient.Get(url)
	if err != nil {
		return VolumesResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			err = errClose
		}
	}(resp.Body)

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return VolumesResponse{}, err
	}

	// Unmarshal the JSON data into the volumesResponse struct
	var volumesResponse VolumesResponse
	err = json.Unmarshal(body, &volumesResponse)
	if err != nil {
		return VolumesResponse{}, err
	}

	// Iterate over the volumes and set the "Type" field to "local" if it's empty
	for i, dockerVolume := range volumesResponse.Volumes {
		if dockerVolume.Type == "" {
			volumesResponse.Volumes[i].Type = "local"
		}
	}

	return volumesResponse, nil
}

// CreateDockerBackupVolume creates a Docker volume where backups are stored
func CreateDockerBackupVolume() (volumeName string, err error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DockerApiUrl))
	if err != nil {
		return "", err
	}
	defer func(cli *client.Client) {
		errClose := cli.Close()
		if errClose != nil {
			err = errClose
		}
	}(cli)

	var volumeConfig volume.CreateOptions
	volumeName = "dvb-backup-" + fmt.Sprint(time.Now().Unix())

	switch config.BackupVolumeType {
	case string(config.Local):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "none",
				"device": config.BackupVolumePath,
				"o":      "bind",
			},
			Labels: map[string]string{
				"dvb.volume.temp": "true",
			},
		}
	case string(config.Nfs):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "nfs",
				"o":      "addr=" + config.BackupVolumeHost + ",rw,noatime,rsize=8192,wsize=8192,tcp,timeo=14,nfsvers=4",
				"device": ":" + config.BackupVolumePath,
			},
			Labels: map[string]string{
				"dvb.volume.temp": "true",
			},
		}
	case string(config.Cifs):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "cifs",
				"o":      "addr=" + config.BackupVolumeHost + ",username=" + config.BackupVolumeUsername + ",password=" + config.BackupVolumePassword + "",
				"device": config.BackupVolumePath,
			},
			Labels: map[string]string{
				"dvb.volume.temp": "true",
			},
		}
	default:
		return "", errors.Errorf("Error while parsing BACKUP_VOLUME_TYPE: %s", config.BackupVolumeType)
	}

	// Create the Docker volume
	_, err = cli.VolumeCreate(context.Background(), volumeConfig)
	if err != nil {
		return volumeName, err
	}

	return volumeName, err
}

// DeleteDockerBackupVolume deletes a Docker volume where backups are stored
// This will not delete data on the disk
func DeleteDockerBackupVolume(volumeName string) error {
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

	// Get information about the Docker volume
	volumeInfo, err := cli.VolumeInspect(context.Background(), volumeName)
	if err != nil {
		return err
	}

	// Check if the volume has the "dvb.volume.temp" label set to "true"
	labelValue, labelExists := volumeInfo.Labels["dvb.volume.temp"]
	if labelExists && strings.ToLower(labelValue) == "true" {
		// If the label is set to "true," delete the volume
		err := cli.VolumeRemove(context.Background(), volumeName, true)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.Errorf("Docker volume %s was not deleted because the label dvb.volume.temp is not set to true.", volumeName)
}
