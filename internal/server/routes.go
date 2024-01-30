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
					FirstName: "Leo",
					LastName:  "Tolstoy",
				},
				{
					ID:        "2",
					FirstName: "Gustave",
					LastName:  "Flaubert",
				},
				{
					ID:        "3",
					FirstName: "F. Scott",
					LastName:  "Fitzgerald",
				},
				{
					ID:        "4",
					FirstName: "William",
					LastName:  "Shakespeare",
				},
				{
					ID:        "5",
					FirstName: "Fyodor",
					LastName:  "Dostoevsky",
				},
			}
			return renderChildren(c, authorview.ListAuthors(authors))
		}
		return renderBase(c, authorview.ListAuthorsWrap())
	})
}
