package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	if os.Getenv("APP_ENV") == "local" {
		logger, _ := zap.NewProduction()
		app.Use(fiberzap.New(fiberzap.Config{
			Logger: logger,
		}))
	}

	//Start the default fiber server
	app.Listen(":8080")
}
