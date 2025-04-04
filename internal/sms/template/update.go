package template

import (
	"strconv"

	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var template domain.SMSTemplate
	var request domain.SMSTemplateUpdateRequest
	repo := hRepository.New(hDb.Get(), &template, c)

	err = repo.GetById(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	updateTemplate(&template, request)

	err = repo.Save()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &template)
}

func updateTemplate(t *domain.SMSTemplate, request domain.SMSTemplateUpdateRequest) {
	if request.Name != "" {
		t.Name = request.Name
	}
	if request.Message != "" {
		t.Message = request.Message
	}
}
