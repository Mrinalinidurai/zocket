package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string
	RabbitMQURL string
	RedisAddr   string
	AWSRegion   string
	S3Bucket    string
}


// LoadConfig loads configuration values from the .env file or environment variables
func LoadConfig() *Config {
	// Load the .env file, if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Return a new Config struct with environment variable values
	return &Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"), // Connection string for PostgreSQL
		RabbitMQURL: os.Getenv("RABBITMQ_URL"), // RabbitMQ connection URL
		RedisAddr:   os.Getenv("REDIS_ADDR"),   // Redis server address
		AWSRegion:   os.Getenv("AWS_REGION"),   // AWS region (e.g., "us-east-1")
		S3Bucket:    os.Getenv("S3_BUCKET"),    // S3 bucket name where files will be stored
	}
}
