package handlers

import (
	dto "service_b/internal/dto"
	service "service_b/internal/service"
	usecase "service_b/internal/usecase"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	otel "go.opentelemetry.io/otel"
	attribute "go.opentelemetry.io/otel/attribute"
	codes "go.opentelemetry.io/otel/codes"
)

var validateCep = validator.New()

func TemperatureRecoveryLocation(c *fiber.Ctx) error {

	ctx := c.UserContext()
	tracer := otel.Tracer("serviceB")

	cep := c.Params("number")

	input := dto.CepDTO{
		Number: cep,
	}

	err := validateCep.Var(input.Number, "required,numeric,len=8")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx, span := tracer.Start(ctx, "Searching for a city by zip code and temperature")
	span.SetAttributes(attribute.String("input.cep", input.Number))

	cepService := service.ServiceCepImpl{}
	weatherApi := service.ServiceWeatherApiImpl{}
	temperatureRecoveryLocationUseCase := usecase.NewTemperatureRecoveryLocationUseCase(
		cepService,
		weatherApi,
	)

	location, err := temperatureRecoveryLocationUseCase.GetTemperatureRecoveryLocation(input.Number)
	if err != nil {
		span.SetStatus(codes.Error, "Error fetching CEP and temperature")
		span.RecordError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, span = tracer.Start(ctx, "Returns zip code location and temperature")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(location)
}
