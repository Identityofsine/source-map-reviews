package config

import (
	"fmt"
	"os"
)

//This file contains the logic that is responsible for details
//of the server, such as the version, branch, commit, and build date.
//This is primarily used in the health component of the web application

type WebServerConfig struct {
	URI string
}

type ServerDetails struct {
	ConfigFile
	ServerName      string `yaml:"server_name"`
	Version         string `yaml:"version"`
	Iteration       string `yaml:"iteration"`
	Commit          string `yaml:"commit"`
	Branch          string `yaml:"branch"`
	Environment     string `yaml:"environment"`
	BuildDate       string //this comes from the environment variable passed to the server on startup
	WebServerConfig WebServerConfig
}

func GetServerDetails() *ServerDetails {
	if config, err := loadConfig[*ServerDetails]("server"); err == nil {
		config.Environment = getEnvironment()
		config.BuildDate = getBuildDate()
		config.WebServerConfig = WebServerConfig{
			URI: os.Getenv("URI"),
		}
		return config
	} else {
		fmt.Printf("Error loading server config: %v\n", err)
		return &ServerDetails{
			ServerName:  "unknown",
			Version:     "x.x.x",
			Iteration:   "-1",
			Commit:      "unknown",
			Branch:      "unknown",
			Environment: getEnvironment(),
			BuildDate:   getBuildDate(),
		}
	}
}

func getEnvironment() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "development"
	}
	return env
}

func getBuildDate() string {
	buildDate := os.Getenv("BUILD_DATE")
	if buildDate == "" {
		return "unknown"
	}
	return buildDate
}
