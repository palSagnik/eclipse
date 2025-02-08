package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/palsagnik2102/eclipse/pkg/docker"
)

func main() {
	dockerService, err := docker.NewDockerService()
	if err != nil {
		fmt.Printf("unable to create docker client: %s\n", err)
	}

	images, err := dockerService.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		fmt.Printf("unable to get images: %s\n", err)
	}

	fmt.Println("Image Details")
	for _, img := range images {
		fmt.Printf("Name: %v \nImage ID: %s \nCreated: %d \n", img.Name, img.ID, img.Created)
	}
}