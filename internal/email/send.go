package email

import (
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"

	emailTemplate "github.com/flambra/sender/internal/email/template"
)

func Send(c *fiber.Ctx) error {
	var request domain.EmailRequest
	var templateEmail domain.TemplateEmail
	repo := hRepository.New(hDb.Get(), &templateEmail, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"name": request.TemplateName})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Template not found")
	}

	body, err := emailTemplate.Process(templateEmail, request.RecipientName)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Failed to process template")
	}

	SMTP := GetSMTPConfig()
	m := gomail.NewMessage()
	m.SetHeader("From", SMTP.From)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", templateEmail.Subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(SMTP.Host, SMTP.Port, SMTP.Username, SMTP.Password)
	
	if err := dialer.DialAndSend(m); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Email sent successfully")
}
