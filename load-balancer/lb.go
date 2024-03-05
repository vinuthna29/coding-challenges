package main

import (
	"fmt"
	"io/ioutil"
	"lb/utils"
	"net/http"
)

type LoadBalancer struct {
	Config *utils.Config
}

func NewLoadBalancer(config *utils.Config) *LoadBalancer {
	return &LoadBalancer{Config: config}
}

func (lb *LoadBalancer) handleRequest(w http.ResponseWriter, r *http.Request) {
	// request recieved on this LB server
	fmt.Println("Received request from:", r.RemoteAddr)

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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error while reading response from backend server", http.StatusInternalServerError)
		return
	}
	fmt.Println("Received response from backend server:", string(body))

	// response to the client
	w.WriteHeader(response.StatusCode)
	w.Write(body)
	fmt.Fprintf(w, "Hello from server!")
	fmt.Println("Sent response to client from load balancer")
}

func (lb *LoadBalancer) Run() {
	http.HandleFunc("/", lb.handleRequest)

	address := fmt.Sprintf("%s:%d", lb.Config.LB_HOST, lb.Config.LB_PORT)
	fmt.Printf("Server listening on %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
