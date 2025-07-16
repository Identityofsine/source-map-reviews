package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config interface {
	Print()
}

type ConfigFile struct {
	Name string
}

func (c *ConfigFile) Print() {
	//get all fields of current struct
	// and print them
	fmt.Printf("Config: %s\n", c.Name)
	fmt.Printf("%+v\n", c)
}

// This file contains the logic that is responsible for details
// of the server, such as the version, branch, commit, and build date.
// This also
func loadConfig[T Config](configName string) (T, error) {
	godotenv.Load(".env")
	var config T
	file, err := os.Open(fmt.Sprintf("config/%s.yaml", configName))
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("Error decoding config file: %v\n", err)
		return config, err
	}

	return config, nil
}
