package usersrepo

import (
	"context"
	"fmt"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/magiconair/properties/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()
	// Run localstack container
	container, err := localstack.RunContainer(ctx, testcontainers.WithImage("localstack/localstack:latest"))
	if err != nil {
		t.Fatal(err)
	}
	// stop and remove localstack container
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	// Testcontainer NewDockerProvider is used to get the provider of the docker daemon
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		t.Fatal("Error in getting docker provider")
	}
	host, err := provider.DaemonHost(ctx)
	if err != nil {
		t.Fatal("Error in getting provider host")
	}
	// Gett external mapped port for the container port
	mappedPort, err := container.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		t.Fatal("Error in getting the external mapped port")
	}
	endpoint := fmt.Sprintf("http://%s:%d", host, mappedPort.Int())
	repo := NewDynamoRepository(endpoint)
	if err := repo.createTable(); err != nil {
		t.Fatal(err)
	}
	newUser := domain.User{
		Id:   "1",
		Name: "Tester",
	}
	gotUser, err := repo.Insert(newUser)
	if err != nil {
		t.Errorf("Got error: %s when inserting new user", err)
	}
	assert.Equal(t, gotUser.Id, "1")
	assert.Equal(t, gotUser.Name, "Tester")

}
