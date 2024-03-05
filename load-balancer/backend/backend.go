package main

import (
	"fmt"
	"lb/utils"
	"net/http"
)

type BackendServer interface {
	handleRequest(w http.ResponseWriter, r *http.Request)
	Run()
}

type BackendServerImpl struct {
	Config *utils.Config
}

func NewBackendServer(config *utils.Config) BackendServer {
	return &BackendServerImpl{Config: config}
}

func (b *BackendServerImpl) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request on backend server:", r.RemoteAddr, r.URL.Path)
	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello From Backend Server\n")
	fmt.Println("Sent response to loadbalancer from backend server")
}

func (b *BackendServerImpl) Run() {
	http.HandleFunc("/", b.handleRequest)

	address := fmt.Sprintf("%s:%d", b.Config.BACKEND_HOST, b.Config.BACKEND_PORT)
	fmt.Printf("Backend server is listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Error starting backend server:", err)
	}
}

func main() {
	logger := utils.NewStdLogger()

	config, err := utils.LoadConfig("../config/config.yaml")
	if err != nil {
		logger.Error("Error loading config:", err)
		return
	}

	bs := NewBackendServer(&config)
	bs.Run()
}
