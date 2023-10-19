package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type VolumeType string

const (
	Nfs   VolumeType = "nfs"
	Cifs  VolumeType = "cifs"
	Local VolumeType = "local"
)

type Volume struct {
	Name      string `json:"Name"`
	UsageData struct {
		Size int64 `json:"Size"`
	} `json:"UsageData"`
	Type   VolumeType        `json:"Type"`
	Labels map[string]string `json:"Labels"`
	//TODO Backup status
}

type VolumesResponse struct {
	Volumes []Volume `json:"Volumes"`
}

// Returns array of all Docker volumes with Name and Type
func GetDockerVolumes() (VolumesResponse, error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DOCKER_API_URL))
	if err != nil {
		return VolumesResponse{}, err
	}
	defer cli.Close()

	volumes, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return VolumesResponse{}, err
	}

	var volumeResponse VolumesResponse
	// Iterate over the list of volumes and get their options
	for _, volume := range volumes.Volumes {
		volumeInfo, err := cli.VolumeInspect(context.Background(), volume.Name)
		if err != nil {
			return VolumesResponse{}, err
		}

		// Convert to json
		opt, err := json.Marshal(volumeInfo.Options)
		if err != nil {
			return VolumesResponse{}, err
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
		volumeResponse.Volumes = append(volumeResponse.Volumes, Volume{
			Name: volume.Name,
			Type: volumeType,
		})
	}

	return volumeResponse, nil
}

// Returns array of all Docker volumes with Name, Size, Type and Labels
func GetDockerVolumesWithSize() (VolumesResponse, error) {
	url := config.DOCKER_API_URL + "/system/df"

	// Send an HTTP  request to the Docker API
	resp, err := http.Get(url)
	if err != nil {
		return VolumesResponse{}, err
	}
	defer resp.Body.Close()

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
	for i, volume := range volumesResponse.Volumes {
		if volume.Type == "" {
			volumesResponse.Volumes[i].Type = "local"
		}
	}

	return volumesResponse, nil
}

// Creates a Docker volume where backups are stored
func CreateDockerBackupVolume() (volumeName string, err error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DOCKER_API_URL))
	if err != nil {
		return "", err
	}
	defer cli.Close()

	var volumeConfig volume.CreateOptions
	volumeName = "dvb-backup-" + fmt.Sprint(time.Now().Unix())
	//var volumeType = config.BACKUP_VOLUME_TYPE

	switch config.BACKUP_VOLUME_TYPE {
	case string(Local):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "none",
				"device": config.BACKUP_VOLUME_PATH,
				"o":      "bind",
			},
			Labels: map[string]string{
				"com.dvb.volume": "true",
			},
		}
	case string(Nfs):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "nfs",
				"o":      "addr=" + config.BACKUP_VOLLUME_HOST + ",rw,noatime,rsize=8192,wsize=8192,tcp,timeo=14,nfsvers=4",
				"device": ":" + config.BACKUP_VOLUME_PATH,
			},
			Labels: map[string]string{
				"com.dvb.volume": "true",
			},
		}
	case string(Cifs):
		volumeConfig = volume.CreateOptions{
			Name:   volumeName,
			Driver: "local",
			DriverOpts: map[string]string{
				"type":   "cifs",
				"o":      "addr=" + config.BACKUP_VOLLUME_HOST + ",username=" + config.BACKUP_VOLUME_USERNAME + ",password=" + config.BACKUP_VOLUME_PASSWORD + "",
				"device": config.BACKUP_VOLUME_PATH,
			},
			Labels: map[string]string{
				"com.dvb.volume": "true",
			},
		}
	default:
		return "", errors.Errorf("Error while parsing BACKUP_VOLUME_TYPE: %s", config.BACKUP_VOLUME_TYPE)
	}

	// Create the Docker volume
	_, err = cli.VolumeCreate(context.Background(), volumeConfig)
	if err != nil {
		return volumeName, err
	}

	return volumeName, err
}

// Deletes a Docker volume where backups are stored
// This will not delete data on the disk
func DeleteDockerBackupVolume(volumeName string) error {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.WithHost(config.DOCKER_API_URL))
	if err != nil {
		return err
	}
	defer cli.Close()

	// Get information about the Docker volume
	volumeInfo, err := cli.VolumeInspect(context.Background(), volumeName)
	if err != nil {
		return err
	}

	// Check if the volume has the "com.dvb.volume" label set to "true"
	labelValue, labelExists := volumeInfo.Labels["com.dvb.volume"]
	if labelExists && strings.ToLower(labelValue) == "true" {
		// If the label is set to "true," delete the volume
		err := cli.VolumeRemove(context.Background(), volumeName, true)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.Errorf("Docker volume %s was not deleted because the label com.dvb.volume is not set to true.", volumeName)
}
