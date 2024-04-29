package main

import (
	v1 "cats-social/internal/delivery/http/v1"
	postgresqlpkg "cats-social/pkg/database/postgresql"
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
	var (
		db = postgresqlpkg.InitPostgreSQL()
	)
	app := fiber.New(fiber.Config{
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: true,
	})

	if os.Getenv("APP_ENV") == "local" {
		logger, _ := zap.NewProduction()
		app.Use(fiberzap.New(fiberzap.Config{
			Logger: logger,
		}))
	}

	// Init Router
	v1.Init(app, db)

	//Start the default fiber server
	app.Listen(":8080")
}
