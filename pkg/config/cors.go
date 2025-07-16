package config

import "fmt"

type CorsConfig struct {
	ConfigFile
	AllowOrigins     []string `yaml:"allow_origins" json:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods" json:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers" json:"allow_headers"`
	ExposeHeaders    []string `yaml:"expose_headers" json:"expose_headers"`
	AllowCredentials bool     `yaml:"allow_credentials" json:"allow_credentials"`
	MaxAge           int      `yaml:"max_age" json:"max_age"`
	AllowAll         bool     `yaml:"allow_all" json:"allow_all"`
}

func (c *CorsConfig) Print() {
	fmt.Printf("CORS Config\n")
	fmt.Printf("AllowOrigins: %v\n", c.AllowOrigins)
	fmt.Printf("AllowMethods: %v\n", c.AllowMethods)
	fmt.Printf("AllowHeaders: %v\n", c.AllowHeaders)
	fmt.Printf("ExposeHeaders: %v\n", c.ExposeHeaders)
	fmt.Printf("AllowCredentials: %v\n", c.AllowCredentials)
	fmt.Printf("MaxAge: %d\n", c.MaxAge)
	fmt.Printf("AllowAll: %v\n", c.AllowAll)
	fmt.Printf("Config: %s\n", c.Name)
}

// GetCorsConfig returns the CORS configuration for the server.
// It loads the configuration from a YAML file and returns a CorsConfig struct.
func GetCorsConfig() *CorsConfig {
	if config, err := loadConfig[*CorsConfig]("cors"); err == nil {
		return config
	} else {
		fmt.Printf("Error loading CORS config: %v\n", err)
		return &CorsConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			ExposeHeaders:    []string{},
			AllowCredentials: false,
			MaxAge:           3601,
			AllowAll:         true,
		}
	}
}
