package controller

import (
	"syscall"
)

// CheckFreeStorageByPath checks if there is enough space on the disk to store the backups. Checks the path directly
func CheckFreeStorageByPath(path string) (bytesFree uint64, err error) {
	var stat syscall.Statfs_t
	err = syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}

	blockSize := stat.Bsize
	blocksAvailable := stat.Bavail

	bytesFree = uint64(blocksAvailable) * uint64(blockSize)

	return bytesFree, nil
}

// CheckFreeStorageByDockerVolume checks if there is enough space on the disk to store the backups. Checks the path used by the docker volume
func CheckFreeStorageByDockerVolume(volumeName string) (bytesFree int, err error) {
	// TODO
	return 0, nil
}

// CheckStorageAvailabilityByPath checks if the selected storage is writable. Checks the path directly
func CheckStorageAvailabilityByPath(path string) (available bool, err error) {
	// TODO
	return false, err
}

// CheckStorageAvailabilityByDockerVolume checks if the selected storage is writable. Checks the path used by the docker volume
func CheckStorageAvailabilityByDockerVolume(volumeName string) (available bool, err error) {
	// TODO
	return false, nil
}
