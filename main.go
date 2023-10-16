package main

import (
	"fmt"

	"github.com/Juli-EXP/docker-volume-backup/docker"
)

func main() {
	volumes, err := docker.GetDockerVolumes()
	if err != nil {
		panic(err)
	}

	for _, volume := range volumes {
		fmt.Println(volume.Name)
		fmt.Println(volume.Type)
	}
}
