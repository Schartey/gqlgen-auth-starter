package main

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTExpireIn time.Duration

	AllowP string

	DebugMode bool
	LogFormat string
}

func LoadConfig(name string, path string) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName(name)
	config.AddConfigPath(path)
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
