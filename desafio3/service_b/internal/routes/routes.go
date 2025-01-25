package routes

import (
	handlers "service_b/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	cepGroup := api.Group("/cep")
	cepGroup.Get("/:number/get", handlers.TemperatureRecoveryLocation)
}
