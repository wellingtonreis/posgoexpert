package handlers

import (
	"service_a/internal/dto"
	"service_a/internal/usecase"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

var validateCep = validator.New()

type CepHandler struct {
	cepUseCase *usecase.CepUseCase
}

func NewCepHandler(uc *usecase.CepUseCase) CepHandler {
	return CepHandler{cepUseCase: uc}
}

func (h *CepHandler) GetCep(c *fiber.Ctx) error {
	var input dto.CepDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := validateCep.Var(input.Number, "required,numeric,len=8"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.cepUseCase.GetCep(input.Number)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(response)
}
