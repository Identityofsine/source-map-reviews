package model

type Health struct {
	ServerName  string `json:"server_name"`
	BuildDate   string `json:"build_date"`
	Version     string `json:"version" yaml:"version"`
	Commit      string `json:"commit" yaml:"commit"`
	Branch      string `json:"branch" yaml:"branch"`
	Environment string `json:"environment" yaml:"environment"`
}
