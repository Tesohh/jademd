package handler

import (
	"jademd/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func TestTempl(c *fiber.Ctx) error {
	v := view.Index("CISSY", view.Tubre())
	hand := adaptor.HTTPHandler(templ.Handler(v))
	return hand(c)
}
