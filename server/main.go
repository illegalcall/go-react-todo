package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/illegalcall/go-react-todo/server/database"
	"github.com/illegalcall/go-react-todo/server/handlers"
)

func main() {
	fmt.Println("🚀 Starting Go Fiber Mongo API...")

	// Initialize Database
	database.InitDB()
	defer func() {
		if err := database.Client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
		fmt.Println("🛑 Disconnected from MongoDB")
	}()

	app := fiber.New()

	// Middlewares
	app.Use(logger.New()) // Logs requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes
	api := app.Group("/api")
	api.Get("/todos", handlers.GetTodos)
	api.Post("/todos", handlers.CreateTodo)
	api.Patch("/todos/:id", handlers.UpdateTodo)
	api.Delete("/todos/:id", handlers.DeleteTodo)

	// Serve frontend if in production
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	// Graceful shutdown
	go func() {
		if err := app.Listen("0.0.0.0:" + port); err != nil {
			log.Fatalf("❌ Server Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("\n💀 Shutting down server...")
	_ = app.Shutdown()
}
