package handlers

import (
	dto "service_a/internal/dto"
	usecase "service_a/internal/usecase"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	otel "go.opentelemetry.io/otel"
	attribute "go.opentelemetry.io/otel/attribute"
	codes "go.opentelemetry.io/otel/codes"
)

var validateCep = validator.New()

type CepHandler struct {
	cepUseCase *usecase.CepUseCase
}

func NewCepHandler(uc *usecase.CepUseCase) CepHandler {
	return CepHandler{cepUseCase: uc}
}

func (h *CepHandler) GetCep(c *fiber.Ctx) error {

	ctx := c.UserContext()
	tracer := otel.Tracer("serviceA")

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

	ctx, span := tracer.Start(ctx, "Searching service B for the ZIP code")
	span.SetAttributes(attribute.String("input.cep", input.Number))

	response, err := h.cepUseCase.GetCep(input.Number)
	if err != nil {
		span.SetStatus(codes.Error, "Error fetching CEP")
		span.RecordError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, span = tracer.Start(ctx, "Returns the response from the zip code and temperature query service")
	defer span.End()

	return c.JSON(response)
}
