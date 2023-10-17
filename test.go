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
	testDeleteVolume("test")
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

func testCreateVolume(){
	volumeName, err := backup.CreateDockerBackupVolume()
	if err != nil {
		panic(err)
	}
	fmt.Println(volumeName)
}

func testDeleteVolume(volumeName string){
	err := backup.DeleteDockerBackupVolume(volumeName)
	if err != nil {
		panic(err)
	}
	testPrintVolumes()
}