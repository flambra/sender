package sms

import (
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/flambra/sender/internal/queue"
	"github.com/gofiber/fiber/v2"

	smsTemplate "github.com/flambra/sender/internal/sms/template"
)

func Send(c *fiber.Ctx) error {
	var request domain.SMSRequest
	var template domain.SMSTemplate
	repo := hRepository.New(hDb.Get(), &template, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"name": request.TemplateName})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Template not found")
	}

	message, err := smsTemplate.Process(template, request.Variables)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Failed to process template")
	}

	queue.EnqueueSMSTask(request, template, message)

	return hResp.SuccessResponse(c, "SMS sent successfully")
}
