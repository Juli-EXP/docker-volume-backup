package config

import (
	"os"
)

var ServerPort = getEnv("SERVER_PORT", "3000")
var DockerApiUrl = getEnv("DOCKER_API_URL", "http://localhost:2375")
var BackupContainerImage = getEnv("BACKUP_CONTAINER_IMAGE", "alpine:latest")
var BackupVolumeType = getEnv("BACKUP_VOLUME_TYPE", "local")
var BackupVolumePath = getEnv("BACKUP_VOLUME_PATH", "/backup")
var BackupVolumeHost = getEnv("BACKUP_VOLUME_HOST", "")
var BackupVolumeUsername = getEnv("BACKUP_VOLUME_USERNAME", "")
var BackupVolumePassword = getEnv("BACKUP_VOLUME_PASSWORD", "")

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

type StorageType string

const (
	Nfs   StorageType = "nfs"
	Cifs  StorageType = "cifs"
	Local StorageType = "local"
)
