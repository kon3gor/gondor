package main

import (
	"fmt"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/kon3gor/gondor"
)

var config struct {
	Clients struct {
		Github ClientConfig `yaml:"github"`
	} `yaml:"clients"`
}

type ClientConfig struct {
	Url     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout"`
}

func CustomTimeUnmarshaler(t *time.Duration, data []byte) error {
	var err error
	*t, err = time.ParseDuration(string(data))
	return err
}

func main() {
	// Using custom unmarshaller for timout
	yaml.RegisterCustomUnmarshaler(CustomTimeUnmarshaler)

	err := gondor.Parse(&config, "./config.yaml", "./config.prod.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
