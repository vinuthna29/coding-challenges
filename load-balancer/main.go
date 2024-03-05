package main

import (
	"lb/utils"
)

func main() {
	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		utils.LogError("Error loading config:", err)
		return
	}

	lb := NewLoadBalancer(config)
	lb.Run()
}
