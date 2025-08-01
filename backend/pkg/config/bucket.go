package config

type BucketConfig struct {
	ConfigFile
	BucketPath string `yaml:"bucket_path" json:"bucket_path"`
}

func GetBucketConfig() *BucketConfig {
	if config, err := loadConfig[*BucketConfig]("bucket"); err == nil {
		return config
	} else {
		return &BucketConfig{}
	}
}
