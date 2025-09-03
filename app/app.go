package app

import (
	"falcon/routers"
	"os"
	"time"

	db "falcon/database"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func AppInstance() *fiber.App {

	if os.Getenv("GO_ENV") != "test" {
		db.Connect()
	} else {
		db.ConnectTestDb()
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024, // request data limit
		AppName:   "falcon",
		Prefork:   os.Getenv("GO_ENV") != "test", // allow multiple process to handle requests | increases throughput
	})

	app.Use(recover.New()) // prevents from server shutdown due to panics
	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Asia/Kolkata",
	})) // writes logs to stdout (only requests including speed)

	if os.Getenv("GO_ENV") != "test" {
		app.Use(swagger.New(swagger.Config{
			BasePath: "/api/swagger",
			FilePath: "./swagger.json",
			Path:     "docs",
			Title:    "Falcon Swagger API Docs",
			CacheAge: 3600,
		})) // swagger configs
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"status":  "success",
			"message": "OK",
		})
	})

	routers.UserRoutesInit(app)

	return app
}
