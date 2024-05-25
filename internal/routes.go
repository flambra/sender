package internal

import (
	"os"

	"github.com/flambra/sender/internal/email"
	"github.com/flambra/sender/internal/email/template"
	"github.com/flambra/sender/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", middleware.Auth, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Post("/email/send", middleware.Auth, email.Send)
	app.Post("/email/template", middleware.Auth, template.Create)
	app.Get("/email/template/:id", middleware.Auth, template.Read)
	app.Put("/email/template/:id", middleware.Auth, template.Update)
	app.Delete("/email/template/:id", middleware.Auth, template.Delete)

}
