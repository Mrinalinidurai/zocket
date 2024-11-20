package main

import (
	"log"
	"product-management/configs"
	"product-management/internal/api"
	"product-management/internal/database"
	"product-management/internal/logging"
	"product-management/internal/queue"
	"product-management/internal/cache"
	"product-management/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	// Initialize services
	logging.InitLogger()
	database.ConnectDB(cfg)
	defer database.CloseDB()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	queue.ConnectRabbitMQ(cfg)
	defer queue.CloseRabbitMQ()
	storage.InitS3(cfg)
	defer storage.CloseS3()

	// Set up router and routes
	router := gin.Default()
	api.RegisterRoutes(router)

	// Start the server
	log.Fatal(router.Run(":8080"))
}
