package main

import (
	"fmt"
	"net/http"
)

const (
	BACKEND_HOST = "127.0.0.1"
	BACKEND_PORT = 9090
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request on backend server:", r.RemoteAddr, r.URL.Path)
	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello From Backend Server\n")
	fmt.Println("Sent response to loadbalancer from backend server")
}

func main() {
	http.HandleFunc("/", handleRequest)

	address := fmt.Sprintf("%s:%d", BACKEND_HOST, BACKEND_PORT)
	fmt.Printf("Backend server is listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Error starting backend server:", err)
	}
}
