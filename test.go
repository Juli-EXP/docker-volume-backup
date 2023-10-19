package main

import (
	"encoding/json"
	"fmt"

	"github.com/Juli-EXP/docker-volume-backup/backup"
)

func main() {
	//printVolumes()
	//printVolumesWithSize()
	//testCreateVolume()
	//testDeleteVolume("test")
	//testCreateVolumeBackup("portainer_data")
	fmt.Println("Hello World")
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testPrintVolumes() {
	volumeResponse, err := backup.GetDockerVolumes()
	if err != nil {
		panic(err)
	}

	for _, volume := range volumeResponse.Volumes {
		fmt.Println(volume.Name)
		fmt.Println(volume.Type)
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testPrintVolumesWithSize() {
	volumeResponse, err := backup.GetDockerVolumesWithSize()
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

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testCreateVolume() {
	volumeName, err := backup.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}
	fmt.Println(volumeName)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testDeleteVolume(volumeName string) {
	err := backup.DeleteDockerBackupVolume(volumeName)
	if err != nil {
		panic(err)
	}
	testPrintVolumes()
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testCreateVolumeBackup(volumeName string) {
	backupVolumeName, err := backup.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}

	err = backup.CreateDockerVolumeBackup(backup.CreateBackupOptions{
		VolumeName:       volumeName,
		BackupVolumeName: backupVolumeName,
		IncludeCifs:      false,
		IncludeNfs:       false,
	})
	if err != nil {
		panic(err)
	}

	backup.DeleteDockerBackupVolume(backupVolumeName)
}
