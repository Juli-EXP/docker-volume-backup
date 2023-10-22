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

func testCreateVolume() {
	volumeName, err := backup.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}
	fmt.Println(volumeName)
}

func testDeleteVolume(volumeName string) {
	err := backup.DeleteDockerBackupVolume(volumeName)
	if err != nil {
		panic(err)
	}
	testPrintVolumes()
}

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

	err = backup.DeleteDockerBackupVolume(backupVolumeName)
	if err != nil {
		panic(err)
	}
}
