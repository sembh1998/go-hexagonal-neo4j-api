package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mooc "github.com/sembh1998/go-hexagonal-api/internal/platform"
	"github.com/sembh1998/go-hexagonal-api/internal/platform/server/handler/products"
	status "github.com/sembh1998/go-hexagonal-api/internal/platform/server/handler/status"
)

type Server struct {
	// httpAddr is the address the server will listen on.
	httpAddr string
	// engine is the HTTP engine used to handle requests.
	engine *gin.Engine

	// productRepository is the repository used to access the products.
	productRepository mooc.ProductRepository
}

// New returns a new Server.
func New(host string, port uint, prodRepo mooc.ProductRepository) *Server {
	server := &Server{
		httpAddr:          fmt.Sprintf("%s:%d", host, port),
		engine:            gin.New(),
		productRepository: prodRepo,
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
	s.engine.POST("/products", products.CreateHandler(s.productRepository))
}
