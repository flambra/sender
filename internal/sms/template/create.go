package template

import (
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var template domain.SMSTemplate
	var request domain.SMSTemplateCreateRequest
	repo := hRepository.New(hDb.Get(), &template, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	if err := repo.GetWhere(fiber.Map{"name": request.Name}); err == nil {
		return hResp.StatusConflict(c, &template, "Name already in use")
	}

	template = domain.SMSTemplate{
		Name:    request.Name,
		Message: request.Message,
	}

	err := repo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, &template)
}
