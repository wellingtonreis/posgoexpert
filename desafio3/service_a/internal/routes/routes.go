package routes

import (
	"log"
	"service_a/internal/di"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	container, err := di.BuildContainerCep()
	if err != nil {
		log.Fatalf("Failed to build container: %v", err)
	}

	api := app.Group("/api/v1")
	cepGroup := api.Group("/cep")
	cepGroup.Post("/post", container.CepHandler.GetCep)
}
