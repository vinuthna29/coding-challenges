package main

import (
	"lb/utils"
)

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
