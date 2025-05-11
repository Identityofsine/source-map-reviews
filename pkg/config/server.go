package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

//This file contains the logic that is responsible for details
//of the server, such as the version, branch, commit, and build date.
//This is primarily used in the health component of the web application

type ServerDetails struct {
	ServerName  string `yaml:"server_name"`
	Version     string `yaml:"version"`
	Iteration   string `yaml:"iteration"`
	Commit      string `yaml:"commit"`
	Branch      string `yaml:"branch"`
	Environment string `yaml:"environment"`
	BuildDate   string //this comes from the environment variable passed to the server on startup
}

func GetServerDetails() *ServerDetails {
	if config, err := loadConfig(); err == nil {
		config.BuildDate = getBuildDate()
		return config
	} else {
		return &ServerDetails{
			ServerName:  "unknown",
			Version:     "x.x.x",
			Iteration:   "-1",
			Commit:      "unknown",
			Branch:      "unknown",
			Environment: "unknown",
			BuildDate:   getBuildDate(),
		}
	}
}

func getBuildDate() string {
	buildDate := os.Getenv("BUILD_DATE")
	if buildDate == "" {
		return "unknown"
	}
	return buildDate
}

func loadConfig() (*ServerDetails, error) {
	config := &ServerDetails{}
	file, err := os.Open("server.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
