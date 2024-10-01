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

	// cissy

	db, err := data.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// pass "globals" to all contexts
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "PublisherKey")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent) // Respond with 204 No Content to preflight
		}

		return c.Next()
	})

	app.Use(logger.New())

	app.Static("/resources", "./view/resources/")
	app.Post("/publish", handler.Publish)
	app.Get("/testtempl", handler.TestTempl)

	app.Listen(":8080")
}
