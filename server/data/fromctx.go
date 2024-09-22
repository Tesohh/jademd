package data

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FromCtx(c *fiber.Ctx) *gorm.DB {
	db, ok := c.Locals("db").(*gorm.DB)
	if !ok {
		log.Fatal("reached unreachable state: `c.Locals(\"db\")` is not of type `*gorm.DB`")
	}
	return db
}
