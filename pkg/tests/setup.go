package tests

import (
	"Go_Food_Delivery/pkg/database"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
	"log/slog"
	"os"
	"time"
)

// Setup will be bootstrapping our test db.
func Setup() database.Database {
	slog.Info("Initializing Setup..")
	testDb := database.NewTestDB()

	if err := testDb.Migrate(); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
	return testDb
}

func Teardown(testDB database.Database) {
	err := testDB.Close()
	if err != nil {
		log.Fatalf("Error closing testDB: %s", err)
	}
	err = os.RemoveAll("./tmp")
	if err != nil {
		log.Fatalf("Error removing ./tmp: %s", err)
	}
}

func CreateNATSContainer(ctx context.Context) string {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	networkName := "nats"

	// Create the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "nats:2.10.20",
		Cmd:   []string{"--http_port", "8222"},
		ExposedPorts: map[nat.Port]struct{}{
			"4222/tcp": {},
			"8222/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			"4222/tcp": {{HostPort: "4222"}},
			"8222/tcp": {{HostPort: "8222"}},
		},
		NetworkMode: container.NetworkMode(networkName),
		AutoRemove:  true,
	}, nil, nil, "nats")

	if err != nil {
		log.Fatal(err)
	}
	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container started, waiting for it to become healthy...")
	time.Sleep(5 * time.Second)
	return resp.ID
}

func RemoveNATSContainer(containerID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	log.Println("Remove::NATS Container")
	if err := cli.ContainerStop(ctx, containerID, container.StopOptions{}); err != nil {
		log.Fatal(err)
	}
}
