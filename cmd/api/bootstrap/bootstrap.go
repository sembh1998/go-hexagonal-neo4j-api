package bootstrap

import "github.com/sembh1998/go-hexagonal-api/internal/platform/server"

const (
	// host is the host name of the server.
	host = "localhost"
	// port is the port number of the server.
	port = 8080
)

// Run starts the server.
func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
