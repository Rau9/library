package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetString(key string) string {
	return viper.GetString(key)
}

func Init() error {
	viper.SetDefault("SERVICE_NAME", "library")
	viper.SetDefault("BIND_ADDRESS", "0.0.0.0:3000")
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_DBNAME", "catalog")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_SSLMODE", "disable")
	viper.SetDefault("LOG_FILE", "/var/log/exported/log.json")
	_ = viper.BindEnv("SERVICE_NAME")
	_ = viper.BindEnv("BIND_ADDRESS")
	_ = viper.BindEnv("ENVIRONMENT")
	_ = viper.BindEnv("POSTGRES_HOST")
	_ = viper.BindEnv("POSTGRES_USER")
	_ = viper.BindEnv("POSTGRES_DBNAME")
	_ = viper.BindEnv("POSTGRES_PORT")
	_ = viper.BindEnv("POSTGRES_SSLMODE")
	_ = viper.BindEnv("POSTGRES_PASSWORD")
	_ = viper.BindEnv("LOG_FILE")
	if GetString("POSTGRES_PASSWORD") == "" {
		return fmt.Errorf("POSTGRES_PASSWORD environment variable is required")
	}
	return nil
}
