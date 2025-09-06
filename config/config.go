package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Storage  StorageConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type StorageConfig struct {
	Endpoint   string
	Bucket     string
	AccessKey  string
	SecretKey  string
	Region     string
	UseSSL     bool
}

type ServerConfig struct {
	Port string
	Host string
}

var AppConfig *Config

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "root"),
			Password: getEnvOrDefault("DB_PASSWORD", "password"),
			Name:     getEnvOrDefault("DB_NAME", "db"),
		},
		Storage: StorageConfig{
			Endpoint:  getEnvOrDefault("MINIO_ENDPOINT", "http://localhost:9000"),
			Bucket:    getEnvOrDefault("MINIO_BUCKET", "content-bucket"),
			AccessKey: getEnvOrDefault("MINIO_ROOT_USER", "minioadmin"),
			SecretKey: getEnvOrDefault("MINIO_ROOT_PASSWORD", "minioadmin123"),
			Region:    getEnvOrDefault("MINIO_REGION", "us-east-1"),
			UseSSL:    getEnvAsBool("MINIO_USE_SSL", false),
		},
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "3000"),
			Host: getEnvOrDefault("SERVER_HOST", "0.0.0.0"),
		},
	}
}

func InitConfig() {
	AppConfig = LoadConfig()
	log.Println("Configuration loaded successfully")
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}