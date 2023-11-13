package main

import (
	"encoding/json"
	"fmt"
	"github.com/Juli-EXP/docker-volume-backup/config"
	"github.com/Juli-EXP/docker-volume-backup/controller"
)

func main() {
	config.BackupVolumePath = "/home/julian/backup"

	//testPrintVolumes()
	testPrintVolumesWithSize()
	//_ = testCreateVolume()
	//testDeleteVolume(volumeName)
	//testCreateVolumeBackup("portainer_data")
	//testCheckFreeStorage()

	fmt.Println("Hello World")

	/* TODO Create Backup
	testPrintVolumesWithSize()
	backupVolumeName, _ := controller.CreateDockerBackupVolume()
	fmt.Printf("Backup volume name: %s\n", backupVolumeName)
	//controller.CheckStorageAvailability
	freeStorage, _ := controller.CheckFreeStorage(controller.StorageLocation{
		Type:     config.Local,
		Path:     "/home/julian/backup",
		Host:     "",
		Username: "",
		Password: ""})
	fmt.Printf("Free storage on location: %d\n", freeStorage)
	//backupSize = -> Calculate size of docker volumes to be backed up
	//if freeStorage <= backupSize -> error

	_ = controller.CreateDockerVolumeBackup(controller.CreateBackupOptions{
		VolumeName:       "portainer_data",
		BackupVolumeName: backupVolumeName,
		IncludeCifs:      false,
		IncludeNfs:       false,
	})
	_ = controller.DeleteDockerBackupVolume(backupVolumeName)
	*/
}
func byteCountBinary(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(size)/float64(div), "KMGTPE"[exp])
}

func testPrintVolumes() {
	volumeResponse, err := controller.GetDockerVolumes()
	if err != nil {
		panic(err)
	}

	for _, volume := range volumeResponse.Volumes {
		fmt.Println(volume.Name)
		fmt.Println(volume.Type)
	}
}

func testPrintVolumesWithSize() {
	volumeResponse, err := controller.GetDockerVolumesWithSize()
	if err != nil {
		panic(err)
	}

	for _, volume := range volumeResponse.Volumes {
		fmt.Println(volume.Name)
		fmt.Println(volume.UsageData.Size)
		fmt.Println(volume.Type)

		labels, _ := json.Marshal(volume.Labels)
		fmt.Println(string(labels))
	}
}

func testCreateVolume() (volumeName string) {
	volumeName, err := controller.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}
	fmt.Println(volumeName)
	return volumeName
}

func testDeleteVolume(volumeName string) {
	err := controller.DeleteDockerBackupVolume(volumeName)
	if err != nil {
		panic(err)
	}
	testPrintVolumes()
}

func testCreateVolumeBackup(volumeName string) {
	backupVolumeName, err := controller.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}

	err = controller.CreateDockerVolumeBackup(controller.CreateBackupOptions{
		VolumeName:       volumeName,
		BackupVolumeName: backupVolumeName,
		IncludeCifs:      false,
		IncludeNfs:       false,
	})
	if err != nil {
		panic(err)
	}

	err = controller.DeleteDockerBackupVolume(backupVolumeName)
	if err != nil {
		panic(err)
	}
}

func testCheckFreeStorage() {
	freeStorage, err := controller.CheckFreeStorage(controller.StorageLocation{
		Type:     config.Local,
		Path:     "/home/julian",
		Host:     "",
		Username: "",
		Password: "",
	})

	if err != nil {
		panic(err)
	}
	freeStorageGB := float64(freeStorage) / (1 << 30)

	fmt.Printf("%d B\n", freeStorage)
	fmt.Printf("%f GB\n", freeStorageGB)
}
