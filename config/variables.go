package config

import (
	"os"
)

var SERVER_PORT = getEnv("SERVER_PORT", "3000")
var DOCKER_API_URL = getEnv("DOCKER_API_URL", "http://localhost:2375")
var BACKUP_CONTAINER_IMAGE = getEnv("BACKUP_CONTAINER_IMAGE", "alpine:latest")
var BACKUP_VOLUME_TYPE = getEnv("BACKUP_VOLUME_TYPE", "local")
var BACKUP_VOLUME_PATH = getEnv("BACKUP_VOLUME_PATH", "/backup")
var BACKUP_VOLLUME_HOST = getEnv("BACKUP_VOLUME_HOST", "")
var BACKUP_VOLUME_USERNAME = getEnv("BACKUP_VOLUME_USERNAME", "")
var BACKUP_VOLUME_PASSWORD = getEnv("BACKUP_VOLUME_PASSWORD", "")

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
