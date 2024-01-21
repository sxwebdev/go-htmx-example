package server

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// this func allows to reload frontend after backend restarted
func (s *Service) dev() {
	s.fiber.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.fiber.Get("/ws", websocket.New(func(c *websocket.Conn) {
		ch := make(chan struct{})

		// block forever
		<-ch
	}))
}
