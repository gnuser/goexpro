package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Proxy   string `yaml: "proxy"`
	Binance struct {
		ApiKey    string `yaml:"api_key"`
		SecretKey string `yaml:"secret_key"`
	}
}

func ParseConfigFile(filename string) *Config {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}
