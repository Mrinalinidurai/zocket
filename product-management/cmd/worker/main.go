package main

import (
	"encoding/json"
	"log"
	"product-management/configs"
	"product-management/internal/database"
	"product-management/internal/logging"
	"product-management/internal/queue"
	"product-management/internal/storage"
)

func main() {
	// Load configurations
	cfg := configs.LoadConfig()

	// Initialize services
	logging.InitLogger()
	database.ConnectDB(cfg)
	defer database.CloseDB()
	queue.ConnectRabbitMQ(cfg) // Connect to RabbitMQ
	defer queue.CloseRabbitMQ()
	storage.InitS3(cfg) // Initialize S3
	defer storage.CloseS3()

	// Retrieve messages from RabbitMQ
	msgs, err := queue.GetRabbitMQMessages()
	if err != nil {
		log.Fatalf("Failed to get messages from RabbitMQ: %v", err)
	}

	// Process the messages
	for msg := range msgs {
		var task map[string]interface{}

		// Parse the message body as JSON
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			msg.Nack(false, false) // Reject the message without re-queueing
			continue
		}

		// Process the task
		processImage(task)

		// Acknowledge the message after successful processing
		if err := msg.Ack(false); err != nil {
			log.Printf("Failed to acknowledge message: %v", err)
			continue
		}

		log.Println("Message processed successfully!")
	}
}

// processImage simulates processing of the image task
func processImage(task map[string]interface{}) {
	// Add your image processing logic here
	log.Printf("Processing task: %v", task)
}
