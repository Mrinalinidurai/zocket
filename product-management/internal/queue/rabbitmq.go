package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go" // Updated import for amqp091-go
	"product-management/configs"
)

var conn *amqp091.Connection
var channel *amqp091.Channel

// ConnectRabbitMQ establishes the connection to RabbitMQ
func ConnectRabbitMQ(cfg *configs.Config) {
	var err error
	conn, err = amqp091.Dial(cfg.RabbitMQURL) // Use amqp091.Dial instead of amqp.Dial
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	// Declare the queue
	_, err = channel.QueueDeclare(
		"image-processing", // Queue name
		true,               // Durable (survives RabbitMQ restart)
		false,              // Auto-delete
		false,              // Exclusive
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	log.Println("Connected to RabbitMQ!")
}

// GetRabbitMQMessages retrieves messages from RabbitMQ
func GetRabbitMQMessages() (<-chan amqp091.Delivery, error) {
	// Declare the queue here as well if needed for the consumer
	msgs, err := channel.Consume(
		"image-processing", // Queue name
		"",                 // Consumer tag (empty means random)
		false,              // Auto-acknowledge
		false,              // Exclusive
		false,              // No-local
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

// PublishToQueue publishes a message to RabbitMQ
func PublishToQueue(message string) error {
	err := channel.Publish(
		"",                      // Exchange (default)
		"image-processing",      // Routing key (Queue name)
		false,                   // Mandatory (if the message cannot be routed, it will be returned)
		false,                   // Immediate (if the message cannot be processed, it will be dropped)
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	return err
}

// CloseRabbitMQ closes the RabbitMQ connection and channel
func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
