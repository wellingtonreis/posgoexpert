package main

import (
	"context"
	"log"
	"service_b/internal/config"
	"service_b/internal/routes"

	zipkin "service_b/internal/pkg/zipkin"

	otelfiber "github.com/gofiber/contrib/otelfiber"
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Use(otelfiber.Middleware())
	tracerProvider := zipkin.SetupOTelMiddleware(app, "serviceB")
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	routes.SetupRoutes(app)
	app.Listen(":9000")
}
