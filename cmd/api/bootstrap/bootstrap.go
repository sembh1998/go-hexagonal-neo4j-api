package bootstrap

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/creating"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/bus/inmemory"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/server"
	neo4jimpl "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/storage/neo4j"
)

const (
	// host is the host name of the server.
	host = "localhost"
	// port is the port number of the server.
	port = 8080

	neo4jHost = "localhost"
	neo4jPort = 7687
	neo4jUser = "neo4j"
	neo4jPass = "tartamudoxd"
)

// Run starts the server.
func Run() error {
	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	if err != nil {
		return err
	}

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	var commandBus = inmemory.NewCommandBus()

	productRepository := neo4jimpl.NewProductRepository(neo4jDriver)
	productService := creating.NewProductService(productRepository)
	createProductCommandHandler := creating.NewCreateProductCommandHandler(productService)
	commandBus.Register(creating.ProductCreateCommandType, createProductCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
