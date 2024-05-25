package email

import (
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

func Send(c *fiber.Ctx) error {
	var template domain.TemplateEmail
	var request domain.EmailRequest
	repo := hRepository.New(hDb.Get(), &template, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"name": request.TemplateName})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Template not found")
	}

	SMTP := GetSMTPConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", SMTP.From)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", template.Body)

	d := gomail.NewDialer(SMTP.Host, SMTP.Port, SMTP.Username, SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Email sent successfully")
}
