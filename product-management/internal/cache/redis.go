package cache

import (
	"encoding/json"
	"log"
	"time"

	"product-management/configs"
	"product-management/internal/models"

	"github.com/go-redis/redis/v8"

	"golang.org/x/net/context"
)

var rdb *redis.Client

func InitRedis(cfg *configs.Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
}

func CacheProduct(id string, product models.Product) string {
	ctx := context.Background()
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Printf("Error marshalling product: %v", err)
		return ""
	}

	// Set the product data in the cache with expiration of 1 hour
	err = rdb.Set(ctx, id, productJSON, time.Hour).Err()
	if err != nil {
		log.Printf("Error caching product: %v", err)
		return ""
	}

	return string(productJSON)
}

func GetProductFromCache(id string) string {
	ctx := context.Background()
	productJSON, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Printf("Error fetching from Redis: %v", err)
		return ""
	}
	return productJSON
}

func CloseRedis() {
	if rdb != nil {
		rdb.Close()
	}
}
