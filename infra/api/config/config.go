package config

import (
	"fmt"
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
	// Try loading .env from current and parent directories
	_ = godotenv.Load()                // Current dir
	_ = godotenv.Load("../.env")       // Parent dir
	_ = godotenv.Load("../../.env")    // Root (if running from infra/api/config)
	_ = godotenv.Load("../../../.env") // Root (if running from deep inside)

	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "postgres"),
			getEnv("DB_NAME", "hrms"),
			getEnv("DB_PORT", "5432"),
		)
	}

	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBName:         getEnv("DB_NAME", "hrms"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBURL:          dbURL,
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
