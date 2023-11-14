package config

import (
	"fmt"
	"os"
)

type StorageType string

const (
	Nfs   StorageType = "nfs"
	Cifs  StorageType = "cifs"
	Local StorageType = "local"
)

// Configuration variables

var ServerPort = getEnv("SERVER_PORT", "3000")
var DockerApiUrl = getEnv("DOCKER_API_URL", "unix:///var/run/docker.sock")
var BackupContainerImage = getEnv("BACKUP_CONTAINER_IMAGE", "alpine:latest")
var BackupVolumeType = getEnv("BACKUP_VOLUME_TYPE", "local")
var BackupVolumePath = getEnv("BACKUP_VOLUME_PATH", "/backup")
var BackupVolumeHost = getEnv("BACKUP_VOLUME_HOST", "")
var BackupVolumeUsername = getEnv("BACKUP_VOLUME_USERNAME", "")
var BackupVolumePassword = getEnv("BACKUP_VOLUME_PASSWORD", "")
var BackupNfsVolume = getEnv("BACKUP_NFS_VOLUME", "false")
var BackupCifsVolume = getEnv("BACKUP_CIFS_VOLUME", "false")

// Information variables

var DvbControllerVersion = "0.0.1"
var DvbServerVersion = "0.0.1"
var DvbClientVersion = "0.0.1"

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func GetAllVariables() {
	fmt.Println()
}
