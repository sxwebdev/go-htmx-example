package server

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sxwebdev/go-htmx-example/internal/views/components"
)

func renderBase(
	c *fiber.Ctx,
	component templ.Component,
	options ...func(*templ.ComponentHandler),
) error {
	componentHandler := templ.Handler(
		components.Base(
			components.Layout(component),
			components.BaseParams{
				Title: getTitle(c),
			},
		),
	)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func renderChildren(
	c *fiber.Ctx,
	component templ.Component,
	options ...func(*templ.ComponentHandler),
) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func getTitle(c *fiber.Ctx) string {
	title, ok := c.Locals("title").(string)
	if !ok || title == "" {
		return "GO + HTMX + Templ + Fiber + Tailwind + Vite + AlpineJS"
	}
	return title
}
