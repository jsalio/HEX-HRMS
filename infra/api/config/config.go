package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database Configuration
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBURL      string

	// Server Configuration
	ServerPort     string
	Environment    string
	JWTSecret      string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBName:         getEnv("DB_NAME", "hrms"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBURL:          getEnv("DATABASE_URL", ""),
		ServerPort:     getEnv("SERVER_PORT", "5000"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		JWTSecret:      getEnv("JWT_SECRET", "default_secret"),
		ReadTimeout:    time.Duration(getEnvInt("READ_TIMEOUT", 10)) * time.Second,
		WriteTimeout:   time.Duration(getEnvInt("WRITE_TIMEOUT", 10)) * time.Second,
		MaxHeaderBytes: getEnvInt("MAX_HEADER_BYTES", 1<<20),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Warning: environment variable %s is not an integer, using fallback %d", key, fallback)
			return fallback
		}
		return i
	}
	return fallback
}
