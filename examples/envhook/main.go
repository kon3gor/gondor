package main

import (
	"fmt"

	"github.com/kon3gor/gondor"
	"github.com/kon3gor/gondor/env"
)

var config struct {
	Postgres struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"postgres"`
}

func main() {
	// Using env hook to parse environment variables
	gondor.RegisterStringHook(env.NewEnvHook())

	err := gondor.Parse(&config, "./config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

}
