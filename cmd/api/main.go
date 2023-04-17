package main

import (
	"log"

	"github.com/sembh1998/go-hexagonal-neo4j-api/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
