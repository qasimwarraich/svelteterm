package dockerclient

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer() string {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to create docker client")
	}

	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Hostname:  "cli-tutor",
			Tty:       true,
			OpenStdin: true,
			Image:     "cli-tutor",
		},
		&container.HostConfig{},
		nil,
		nil,
		"")
	if err != nil {
		log.Fatal("Unable to create docker container")
	}

	log.Printf("Container %s created", cont.ID)
	return cont.ID
}

func RunContainer(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to create docker client")
		log.Fatal(err)
	}

	if err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Unable to start container %s", id)
		log.Fatal(err)
	}
	log.Printf("Container %s started", id)
}

func StopContainer(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to create docker client")
		log.Fatal(err)
	}

	if err := cli.ContainerStop(context.Background(), id, nil); err != nil {
		log.Fatalf("Unable to stop container %s", id)
		log.Fatal(err)
	}
	log.Printf("Container %s stopped", id)
}
