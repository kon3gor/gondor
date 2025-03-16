package main

import (
	"flag"
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

var patch string

func main() {
	flag.StringVar(
		&patch,
		"patch",
		"",
		"Provides a 'patch' for the config that will be used as a final layer",
	)

	flag.Parse()

	layers := make([]string, 0, 2)

	// Applying a proper env depended layer
	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "testing"
	}

	if env == "production" {
		layers = append(layers, "./config.prod.yaml")
	} else {
		layers = append(layers, "./config.test.yaml")
	}

	// Add patch if is exists
	if patch != "" {
		layers = append(layers, patch)
	}

	err := gondor.Parse(&config, "./config.yaml", layers...)
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

}
