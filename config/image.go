package config

import "os"

type ImageConfig struct {
	LocalStoragePath string
}

func NewImageConfig() *ImageConfig {
	return &ImageConfig{
		LocalStoragePath: os.Getenv("LOCAL_STORAGE_PATH"),
	}
}
