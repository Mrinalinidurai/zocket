package models

// ProductRequest represents the incoming data for creating a new product
type ProductRequest struct {
	UserID             int      `json:"user_id"` // Change to int instead of string
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	ProductPrice       float64  `json:"product_price"`
}

// Product represents a product in the database
type Product struct {
	ID                      int      `json:"id"`
	UserID                  string   `json:"user_id"`
	ProductName             string   `json:"product_name"`
	ProductDescription      string   `json:"product_description"`
	ProductImages           []string `json:"product_images"`
	CompressedProductImages []string `json:"compressed_product_images"`
	ProductPrice            float64  `json:"product_price"`
	CreatedAt               string   `json:"created_at"`
}
