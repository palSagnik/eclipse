package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/palsagnik2102/eclipse/pkg/docker"
	"github.com/palsagnik2102/eclipse/pkg/tui"
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

	app := tui.NewApp()
	if err := app.ListImage(images...); err != nil {
		fmt.Println(err)
	}
}