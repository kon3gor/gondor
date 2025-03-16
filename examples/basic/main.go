package main

import (
	"fmt"
	"os"

	"github.com/kon3gor/gondor"
)

var config struct {
	Clients struct {
		Github struct {
			Url     string `yaml:"url"`
			Metrics struct {
				Enable bool `yaml:"enable"`
			} `yaml:"metrics"`
		} `yaml:"github"`
	} `yaml:"clients"`
}

func main() {
	// Using just base layer
	err := gondor.Parse(&config, "./config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	// Applying a proper env depended layer
	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "testing"
	}

	if env == "production" {
		err = gondor.Parse(&config, "./config.yaml", "./config.prod.yaml")
	} else {
		err = gondor.Parse(&config, "./config.yaml", "./config.test.yaml")
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
