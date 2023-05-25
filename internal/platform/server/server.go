package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/server/handler/products"
	status "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/server/handler/status"
	"github.com/sembh1998/go-hexagonal-neo4j-api/kit/command"
)

type Server struct {
	// httpAddr is the address the server will listen on.
	httpAddr string
	// engine is the HTTP engine used to handle requests.
	engine *gin.Engine

	commandBus command.Bus
}

// New returns a new Server.
func New(host string, port uint, commandBus command.Bus) *Server {
	server := &Server{
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		engine:     gin.New(),
		commandBus: commandBus,
	}

	server.engine.Use(gin.Recovery())

	server.registerRoutes()

	return server
}

// Run starts the server.
func (s *Server) Run() error {
	return s.engine.Run(s.httpAddr)
}

// registerRoutes registers the routes for the server.
func (s *Server) registerRoutes() {
	s.engine.GET("/status", status.StatusHandler())
	s.engine.POST("/products", products.CreateHandler(s.commandBus))
}
