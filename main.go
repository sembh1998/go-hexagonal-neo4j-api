package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Starting server on", httpAddr)

	mux := http.NewServeMux()

	mux.HandleFunc("/", helloworld)

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request from", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
