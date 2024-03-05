package main

import (
	"fmt"
	"net/http"
)

const (
	LB_HOST = "127.0.0.1"
	LB_PORT = 8080
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from:", r.RemoteAddr) // request recieved on this LB server
	fmt.Fprintf(w, "Hello from server!") // response to the client
}

func main() {
	http.HandleFunc("/", handleRequest)

	address := fmt.Sprintf("%s:%d", LB_HOST, LB_PORT)
	fmt.Printf("Server listening on %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
