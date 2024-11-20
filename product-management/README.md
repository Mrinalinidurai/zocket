 Product Management System

This repository contains a Product Management System built with a microservice-style architecture in Go. The system includes APIs for managing products, a worker service for image processing, caching with Redis, queuing with RabbitMQ, storage with AWS S3, and a PostgreSQL database.



 Features

1. Product APIs:
   - Create, fetch, and list products.
   - Cache frequently accessed products in Redis for improved performance.

2. Worker Service:
   - Processes images for products (e.g., compression).
   - Stores processed images in AWS S3.
   - Acknowledges tasks in RabbitMQ after processing.

3. Integration:
   - PostgreSQL for relational data storage.
   - RabbitMQ for message queuing between APIs and workers.
   - Redis for caching product details.
   - AWS S3 for storing images.

4. Structured Logging:
   - Uses Logrus for detailed JSON-structured logs.



 Project Structure


.
├── cmd
│   ├── api
│   │   └── main.go       # Entry point for the API service
│   ├── worker
│       └── main.go       # Entry point for the worker service
├── configs
│   └── config.go         # Configuration loader for environment variables
├── internal
│   ├── api
│   │   ├── product_handler.go  # API logic for managing products
│   │   └── routes.go     # Routes for API endpoints
│   ├── cache
│   │   └── redis.go      # Redis connection and caching functions
│   ├── database
│   │   └── postgres.go   # PostgreSQL connection management
│   ├── logging
│   │   └── logger.go     # Logging setup and helpers
│   ├── models
│   │   └── product.go    # Data models for products
│   ├── queue
│   │   └── rabbitmq.go   # RabbitMQ connection and queue helpers
│   └── storage
│       └── s3.go         # AWS S3 connection and storage helpers



 Prerequisites

Ensure you have the following installed:
- Go: Version 1.18 or later
- PostgreSQL: For relational database storage
- RabbitMQ: For message queuing
- Redis: For caching
- AWS CLI: For managing S3 buckets
- Docker (optional): To simplify running services locally

---

 Setup

1. Create folder:
   mkdir product-management
   cd product-management
  

2. Create a `.env` file in the root directory and populate it with:
   
   POSTGRES_DSN=your_postgres_dsn
   RABBITMQ_URL=your_rabbitmq_url
   REDIS_ADDR=your_redis_address
   AWS_REGION=your_aws_region
   S3_BUCKET=your_s3_bucket
  

3. Run PostgreSQL, RabbitMQ, and Redis locally or connect to hosted services.

4. Install dependencies:
   
   go mod tidy


   
 Running the Services

 API Service
Start the API service to manage products:

go run cmd/api/main.go

 Worker Service
Start the worker service to process tasks from RabbitMQ:

go run cmd/worker/main.go



 API Endpoints

| Method | Endpoint               | Description               |
|--------|------------------------|---------------------------|
| POST   | `/products`            | Create a new product      |
| GET    | `/products/:id`        | Fetch a product by ID     |
| GET    | `/products`            | List all products         |



 Technologies Used

- Gin: HTTP web framework for Go
- PostgreSQL: Relational database for structured data
- Redis: In-memory data store for caching
- RabbitMQ: Message broker for task queuing
- AWS S3: Object storage for images
- Logrus: Logging with JSON formatting



Author

Created by Mrinalini.

