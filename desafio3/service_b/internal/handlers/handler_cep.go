package handlers

import (
	dto "service_b/internal/dto"
	service "service_b/internal/service"
	usecase "service_b/internal/usecase"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

var validateCep = validator.New()

func TemperatureRecoveryLocation(c *fiber.Ctx) error {
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

	cepService := service.ServiceCepImpl{}
	weatherApi := service.ServiceWeatherApiImpl{}
	temperatureRecoveryLocationUseCase := usecase.NewTemperatureRecoveryLocationUseCase(
		cepService,
		weatherApi,
	)

	location, err := temperatureRecoveryLocationUseCase.GetTemperatureRecoveryLocation(input.Number)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(location)
}
