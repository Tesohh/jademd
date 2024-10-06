package handler

import (
	"fmt"
	"jademd/obsidian"
	"jademd/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func CoursePage(c *fiber.Ctx) error {
	vault, err := obsidian.VaultFromLatest(true, false, false)
	if err != nil {
		return fmt.Errorf("error during vault reading: %w", err)
	}

	// TODO: add enrolled course checking
	// TODO: add markdown view of vault
	v := view.Index(vault.Name, view.CoursePage(vault.Name, vault.Courses, []obsidian.Course{}))
	hand := adaptor.HTTPHandler(templ.Handler(v))
	return hand(c)
}
