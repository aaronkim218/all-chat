package server

import (
	"go-chat/internal/handlers"
	"go-chat/internal/hub"
	"go-chat/internal/storage"
	"go-chat/internal/xerrors"
	"net/http"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Config struct {
	Storage storage.Storage
	Hub     *hub.Hub
}

func New(cfg *Config) *fiber.App {
	app := createFiberApp()
	setupMiddleware(app)
	setupHealthcheck(app)

	service := handlers.NewService(&handlers.ServiceConfig{
		Storage: cfg.Storage,
		Hub:     cfg.Hub,
	})
	service.RegisterRoutes(app)

	setupStatic(app)

	return app
}

func createFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: xerrors.ErrorHandler,
	})
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
}

func setupHealthcheck(app *fiber.App) {
	app.Get("/api/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}

func setupStatic(app *fiber.App) {
	app.Static("/", "internal/static")
}
