package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/palsagnik2102/eclipse/pkg/commands"
)

type DockerService struct {
	cli *client.Client
}

func NewDockerService() (*DockerService, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		// TODO(sagnik): Implement logger
		return nil, err
	}

	return &DockerService{
		cli: cli,
	}, nil
}

func (ds *DockerService) ImageList(ctx context.Context, imgOptions image.ListOptions) ([]commands.Image, error) {
	imageList, err := ds.cli.ImageList(ctx, imgOptions)
	if err != nil {
		// TODO(sagnik): Implement Logger
		return nil, err
	}

	var images []commands.Image
	for _, img := range imageList {
		images = append(images, commands.Image{
			Name: img.RepoTags[0], // TODO(sagnik): Refactor this
			ID: img.ID,
			Created: img.Created,
		})
	}

	return images, nil
}