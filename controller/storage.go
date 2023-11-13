package controller

import (
	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/pkg/errors"
	"syscall"
)

type StorageLocation struct {
	Type     config.StorageType
	Path     string
	Host     string // Only NFS/CIFS
	Username string // Only CIFS
	Password string // Only CIFS
}

// CheckFreeStorage checks if there is enough space on the disk to store the backups
// TODO change so it checks the controller volume, remove StorageLocation struct and add volume name as parameter
func CheckFreeStorage(location StorageLocation) (freeStorage int, err error) {
	switch location.Type {
	case config.Local:
		freeStorage, err = checkFreeStorageLocal(location)
	case config.Nfs:
		freeStorage, err = checkFreeStorageNfs(location)
	case config.Cifs:
		freeStorage, err = checkFreeStorageCifs(location)
	default:
		err = errors.Errorf("No StorageType was specified")
	}

	return freeStorage, err
}

// CheckStorageAvailability checks if the selected storage is writable
func CheckStorageAvailability(location StorageLocation) (available string, err error) {
	return "", err
}

// TODO check with the help of a mounted docker volume
// checkFreeStorageLocal return free storage of a local mount
func checkFreeStorageLocal(location StorageLocation) (freeStorage int, err error) {
	var stat syscall.Statfs_t

	if err = syscall.Statfs(location.Path, &stat); err != nil {
		return 0, err
	}

	freeStorage = int(stat.Bavail * uint64(stat.Bsize))

	return freeStorage, err
}

// TODO check with the help of a mounted docker volume
// checkFreeStorageNfs return free storage of a nfs mount
func checkFreeStorageNfs(location StorageLocation) (freeStorage int, err error) {

	return 0, err
}

// TODO check with the help of a mounted docker volume
// checkFreeStorageCifs return free storage of a cifs mount
func checkFreeStorageCifs(location StorageLocation) (freeStorage int, err error) {

	return 0, err
}
