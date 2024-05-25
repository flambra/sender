package internal

import (
	"os"

	"github.com/flambra/sender/internal/email"
	"github.com/flambra/sender/internal/middleware"
	"github.com/flambra/sender/internal/sms"
	"github.com/gofiber/fiber/v2"

	emailTemplate "github.com/flambra/sender/internal/email/template"
	smsTemplate "github.com/flambra/sender/internal/sms/template"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", middleware.Auth, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	// EMAIL
	app.Post("/email/send", middleware.Auth, email.Send)
	app.Post("/email/template", middleware.Auth, emailTemplate.Create)
	app.Get("/email/template/:id", middleware.Auth, emailTemplate.Read)
	app.Put("/email/template/:id", middleware.Auth, emailTemplate.Update)
	app.Delete("/email/template/:id", middleware.Auth, emailTemplate.Delete)

	// SMS
	app.Post("/sms/send", middleware.Auth, sms.Send)
	app.Post("/sms/template", middleware.Auth, smsTemplate.Create)
	app.Get("/sms/template/:id", middleware.Auth, smsTemplate.Read)
	app.Put("/sms/template/:id", middleware.Auth, smsTemplate.Update)
	app.Delete("/sms/template/:id", middleware.Auth, smsTemplate.Delete)
}
