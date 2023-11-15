package database

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	testcontainers.Container
	Config PostgresContainerConfig
}

type PostgresContainerOption func(c *PostgresContainerConfig)

type PostgresContainerConfig struct {
	DBConfig
	ImageTag string
}

func NewPostgresContainer(ctx context.Context, opts ...PostgresContainerOption) (*PostgresContainer, error) {
	const (
		psqlImage = "postgres"
		psqlPort  = 5432
	)

	containerPort := fmt.Sprintf("%d/tcp", psqlPort)

	config := PostgresContainerConfig{
		ImageTag: "16.1-alpine3.18",
		DBConfig: DBConfig{
			User:     "user",
			Password: "password",
			Database: "database_test",
		},
	}

	// handle possible options
	for _, opt := range opts {
		opt(&config)
	}

	container, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Env: map[string]string{
					"POSTGRES_USER":     config.User,
					"POSTGRES_PASSWORD": config.Password,
					"POSTGRES_DB":       config.Database,
				},
				ExposedPorts: []string{containerPort},
				Image:        fmt.Sprintf("%s:%s", psqlImage, config.ImageTag),
				WaitingFor:   wait.ForListeningPort(nat.Port(containerPort)),
			},
			Started: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("getting request provider: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting host for: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(containerPort))
	if err != nil {
		return nil, fmt.Errorf("getting mapped port for (%s): %w", containerPort, err)
	}

	config.Host = host
	config.Port = mappedPort.Port()

	fmt.Println("Host:", config.Host, config.Port)

	return &PostgresContainer{
		Container: container,
		Config:    config,
	}, nil
}
