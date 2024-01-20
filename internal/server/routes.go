package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sxwebdev/go-htmx-example/internal/models"
	"github.com/sxwebdev/go-htmx-example/internal/views/authorview"
)

func (s *Service) applyRoutes() {
	s.fiber.Get("/", func(c *fiber.Ctx) error {
		if isHtmxRequest(c) {
			authors := []models.Author{
				{
					ID:        "1",
					FirstName: "Alexander",
					LastName:  "Pushkin",
				},
				{
					ID:        "2",
					FirstName: "Mikhail",
					LastName:  "Lermontov",
				},
				{
					ID:        "3",
					FirstName: "Lev",
					LastName:  "Tolstoy",
				},
			}
			return renderChildren(c, authorview.ListAuthors(authors))
		}
		return renderBase(c, authorview.ListAuthorsWrap())
	})
}
