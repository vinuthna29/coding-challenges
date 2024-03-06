package main

import (
	"encoding/json"
	"fmt"
	"lb/utils"
	"net/http"
)

type LoadBalancer interface {
	handleRequest(w http.ResponseWriter, r *http.Request)
	Run()
}
type LoadBalancerImpl struct {
	Config *utils.Config
	// BackendResponse BackendResponse
}

type BackendResponse struct {
	RequestDetails string
	ResponseBody   string
}

func NewLoadBalancer(config *utils.Config) LoadBalancer {
	return &LoadBalancerImpl{Config: config}
}

func (lb *LoadBalancerImpl) handleRequest(w http.ResponseWriter, r *http.Request) {
	// request recieved on this LB server
	// display those details
	// fmt.Println("Received request from:", r.RemoteAddr)
	fmt.Printf("Received request from %s \n", r.RemoteAddr)
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	fmt.Println("Host:", r.Host)
	fmt.Println("Headers:")
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("\t%s: %s\n", name, h)
		}
	}
	fmt.Println("")

	// forward the recieved request to the backend server
	fmt.Println("Forwarding request to backend server:", r.URL.Path)

	backendUrl := fmt.Sprintf("http://%s:%d%s", lb.Config.BACKEND_HOST, lb.Config.BACKEND_PORT, r.URL.Path)
	response, err := http.Get(backendUrl)
	if err != nil {
		http.Error(w, "Error forwarding request to backend server", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// read response from backend server
	var backendResponse BackendResponse
	if err := json.NewDecoder(response.Body).Decode(&backendResponse); err != nil {
		http.Error(w, "Error decoding backend response", http.StatusInternalServerError)
		return
	}

	fmt.Println("Received response from backend server:", backendResponse.ResponseBody)
	// response to the client

	fmt.Println(backendResponse.RequestDetails)

	w.WriteHeader(response.StatusCode)
	w.Write([]byte(backendResponse.ResponseBody))
	fmt.Fprintf(w, "Hello from server!")
	fmt.Println("Sent response to client from load balancer")
}

func (lb *LoadBalancerImpl) Run() {
	http.HandleFunc("/", lb.handleRequest)

	address := fmt.Sprintf("%s:%d", lb.Config.LB_HOST, lb.Config.LB_PORT)
	fmt.Printf("Server listening on %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	logger := utils.NewStdLogger()

	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Error("Error loading config:", err)
		return
	}

	lb := NewLoadBalancer(&config)
	lb.Run()
}
