package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/adapter/handler"

	"github.com/shopally-ai/internal/adapter/gateway"

	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
	"github.com/shopally-ai/pkg/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to MongoDB using custom db package
	client, err := platform.Connect(cfg.Mongo.URI)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := platform.Disconnect(client); err != nil {
			log.Printf("failed to disconnect MongoDB: %v", err)
		}
	}()
	db := client.Database(cfg.Mongo.Database)

	fmt.Printf("Connected to MongoDB database: %s\n", db.Name())

	// Initialize Redis client
	rdb := platform.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx); err != nil {
		log.Printf("⚠️  Redis connection failed: %v (continuing without Redis)", err)
		rdb = nil
	} else {
		log.Println("✅ Redis connected")
	}

	// Initialize router
	router := gin.Default()

	// Construct mock gateways and use case for mocked search flow
	ag := gateway.NewMockAlibabaGateway()
	lg := gateway.NewMockLLMGateway()
	uc := usecase.NewSearchProductsUseCase(ag, lg, nil)

	// Initialize handlers
	searchHandler := handler.NewSearchHandler(uc)

	// Register routes
	searchHandler.RegisterRoutes(router)

	// Start the server
	log.Println("Starting server on port", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
