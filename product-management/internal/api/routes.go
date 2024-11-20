package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/products", CreateProduct)
	router.GET("/products/:id", GetProduct)
	router.GET("/products", ListProducts)
}
