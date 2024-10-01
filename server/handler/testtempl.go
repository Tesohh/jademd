package handler

import (
	"jademd/obsidian"
	"jademd/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func TestTempl(c *fiber.Ctx) error {
	v := view.Index("CISSY", view.Tubre([]obsidian.Course{
		{
			Name:           "Masimelian",
			CourseMetadata: obsidian.CourseMetadata{Color: "red"},
		},
		{
			Name:           "Tubre",
			CourseMetadata: obsidian.CourseMetadata{Color: "blue"},
		},
		{
			Name:           "very long title yes very long super long title",
			CourseMetadata: obsidian.CourseMetadata{Color: "red"},
		},
		{
			Name:           "How to tubre a tauferrer",
			CourseMetadata: obsidian.CourseMetadata{Color: "blue"},
		},
	}))
	hand := adaptor.HTTPHandler(templ.Handler(v))
	return hand(c)
}
