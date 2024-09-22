package main

import (
	"fmt"
	"jademd/data"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

// @title Jade Server
// @version 1.0
// @description Server for your obsidian vault
// @contact.name Simone Tesini
// @contact.url https://github.com/Tesohh
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	fmt.Println("connecting to db...")
	db, err := data.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db, err
	fmt.Println("done")

	app.Get("/", func(c *fiber.Ctx) error {
		c.JSON("CISSYGOZOZFKLJioewfioj")
		return nil
	})

	app.Listen(":8080")
}
