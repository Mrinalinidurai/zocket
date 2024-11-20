package api

import (
	"encoding/json"
	"log"
	"net/http"
	"product-management/internal/cache"
	"product-management/internal/database"
	"product-management/internal/models"
	"product-management/internal/queue"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func CreateProduct(c *gin.Context) {
	var req models.ProductRequest

	// Log the incoming JSON request for debugging purposes
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err) // Log the error for more details
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Received Product Request: %+v", req) // Log the received request for debugging

	db := database.GetDB()

	var productID int
	err := db.QueryRow(`
        INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
        VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		req.UserID, req.ProductName, req.ProductDescription, pq.Array(req.ProductImages), req.ProductPrice,
	).Scan(&productID)

	if err != nil {
		log.Printf("Error inserting into database: %v", err) // Log any database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
		return
	}

	task := map[string]interface{}{
		"product_id": productID,
		"image_urls": req.ProductImages,
	}
	taskJSON, _ := json.Marshal(task)
	if err := queue.PublishToQueue(string(taskJSON)); err != nil {
		log.Printf("Error publishing to queue: %v", err) // Log any queue errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product_id": productID})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	productJSON := cache.GetProductFromCache(id)
	if productJSON != "" {
		c.String(http.StatusOK, productJSON)
		return
	}

	db := database.GetDB()
	var product models.Product
	err := db.QueryRow(`
		SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at
		FROM products WHERE id = $1`, id,
	).Scan(
		&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription,
		(*pq.StringArray)(&product.ProductImages), (*pq.StringArray)(&product.CompressedProductImages),
		&product.ProductPrice, &product.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	productJSON = cache.CacheProduct(id, product)
	c.String(http.StatusOK, productJSON)
}

func ListProducts(c *gin.Context) {
	// Similar to GetProduct but fetch multiple products
}
