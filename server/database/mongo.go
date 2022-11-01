package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
	once       sync.Once
)

// InitDB initializes MongoDB connection
func InitDB() {
	once.Do(func() {
		if os.Getenv("ENV") != "production" {
			if err := godotenv.Load(".env"); err != nil {
				log.Fatal("Error loading .env file:", err)
			}
		}

		uri := os.Getenv("MONGODB_URI")
		clientOptions := options.Client().ApplyURI(uri)

		var err error
		Client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal("MongoDB connection error:", err)
		}

		if err = Client.Ping(context.Background(), nil); err != nil {
			log.Fatal("MongoDB ping error:", err)
		}

		Collection = Client.Database("golang_db").Collection("todos")
		fmt.Println("✅ Connected to MongoDB!")
	})
}
