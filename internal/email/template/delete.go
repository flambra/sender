package template

import (
	"strconv"

	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var template domain.TemplateEmail
	repo := hRepository.New(hDb.Get(), &template, c)

	err = repo.Delete(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = repo.GetDeleted(fiber.Map{"id": id})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "template deleted successfully")
}
