package main

import (
	"jademd/data"
	"log"

	"github.com/gofiber/fiber/v2"
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

	app.Listen(":8080")
}
