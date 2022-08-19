package dockerclient

import (
	"context"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer() string {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print("Unable to create docker client")
	}

	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Hostname:  "cli-tutor",
			Tty:       true,
			OpenStdin: true,
			Image:     os.Getenv("DOCKERIMAGE"),
		},
		&container.HostConfig{},
		nil,
		nil,
		"")
	if err != nil {
		log.Print("Unable to create docker container")
	}

	log.Printf("Container %s created", cont.ID)
	return cont.ID
}

func RunContainer(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print("Unable to create docker client")
		log.Print(err)
	}

	if err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		log.Printf("Unable to start container %s", id)
		log.Print(err)
	}
	log.Printf("Container %s started", id)
}

func StopContainer(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print("Unable to create docker client")
		log.Print(err)
	}

	if err := cli.ContainerStop(context.Background(), id, nil); err != nil {
		log.Printf("Unable to stop container %s", id)
		log.Print(err)
	}
	log.Printf("Container %s stopped", id)
}

func HandlePty(id string, height uint, width uint) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print("Unable to create docker client")
		log.Print(err)
	}

	if err := cli.ContainerResize(context.Background(), id, types.ResizeOptions{
		Height: height,
		Width:  width,
	}); err != nil {
		log.Printf("Unable to resize container %s", id)
		log.Print(err)
	}
	log.Printf("Container %s resized to %d %d", id, height, width)
}
