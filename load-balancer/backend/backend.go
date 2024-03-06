package main

import (
	"encoding/json"
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

type BackendResponse struct {
	RequestDetails string
	ResponseBody   string
}

func NewBackendServer(config *utils.Config) BackendServer {
	return &BackendServerImpl{Config: config}
}

func (b *BackendServerImpl) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request on backend server:", r.RemoteAddr, r.URL.Path)
	// w.WriteHeader(http.StatusOK)
	requestDetails := fmt.Sprintf("Received request from %s \n", r.RemoteAddr)
	requestDetails += fmt.Sprintf("%s %s %s\n", r.Method, r.URL, r.Proto)
	requestDetails += "Headers:\n"
	for name, headers := range r.Header {
		for _, h := range headers {
			requestDetails += fmt.Sprintf("\t%s: %s\n", name, h)
		}
	}
	requestDetails += fmt.Sprintf("Host: %s\n", r.Host)

	fmt.Println(requestDetails)
	fmt.Fprintf(w, "Hello From Backend Server\n")

	backendResponse := BackendResponse{
		RequestDetails: requestDetails,
		ResponseBody:   "Hello From Backend Server\n",
	}

	jsonResponse, err := json.Marshal(backendResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

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
