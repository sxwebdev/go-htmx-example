package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sxwebdev/go-htmx-example/frontend"
	"github.com/sxwebdev/go-htmx-example/internal/config"
	"github.com/sxwebdev/go-htmx-example/internal/views/components"
	"github.com/tkcrm/mx/logger"
	"github.com/tkcrm/mx/service"
)

const serviceName = "frontend-server"

type Service struct {
	logger logger.Logger
	config *config.Config
	name   string

	fiber *fiber.App
}

func New(l logger.ExtendedLogger, cfg *config.Config) *Service {
	s := &Service{
		logger: logger.With(l, "service", serviceName),
		config: cfg,
		name:   serviceName,
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	s.fiber = app

	// add recover
	app.Use(recoverMiddleware.New(recoverMiddleware.Config{
		EnableStackTrace: true,
	}))

	// add cors config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodHead,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// apply routes
	s.applyRoutes()

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontend.StaticFS),
		PathPrefix: "dist",
	}))

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(frontend.StaticFS),
		PathPrefix: "dist/assets",
	}))

	app.Use("*", func(c *fiber.Ctx) error {
		c.Locals("title", "Page not found")
		return renderBase(c, components.NotFound())
	})

	return s
}

func (s Service) Name() string { return s.name }

func (s *Service) Start(ctx context.Context) error {
	errChan := make(chan error, 1)

	// start http server
	go func() {
		errChan <- s.fiber.Listen(s.config.StaticServerAddr)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
	}

	return nil
}

func (s Service) Stop(ctx context.Context) error {
	if s.fiber == nil {
		return nil
	}
	return s.fiber.ShutdownWithContext(ctx)
}

var _ service.IService = (*Service)(nil)
