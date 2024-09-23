package main

import (
	"jademd/data"
	"jademd/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// pass "globals" to all contexts
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
	app.Use(logger.New())

	app.Post("/publish", handler.Publish)

	app.Listen(":8080")
}
