package utils

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	LB_HOST      string `yaml."lb_host"`
	LB_PORT      int    `yaml."lb_port"`
	BACKEND_HOST string `yaml."backend_host"`
	BACKEND_PORT int    `yaml."backend_port"`
}

func LoadConfig(filepath string) (Config, error) {
	var config Config

	// reading data
	data, err := ioutil.ReadFile("../config/config.yaml")
	if err!=nil{
		return config, err
	}

	// decoding data
	err = yaml.Unmarshal(data, &config)
	if err!=nil{
		return config, err
	}
	return config, nil
}
