package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "service_b/internal/handlers"
	routes "service_b/internal/routes"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/v1/cep/12345/get", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTemperatureRecoveryLocation(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/cep/:number/get", handlers.TemperatureRecoveryLocation)

	req := httptest.NewRequest("GET", "/api/v1/cep/12345/get", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
