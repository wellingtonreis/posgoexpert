package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"service_b/internal/handlers"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureRecoveryLocation_ValidCep(t *testing.T) {
	app := fiber.New()
	app.Get("/temperature/:number", handlers.TemperatureRecoveryLocation)

	req := httptest.NewRequest("GET", "/temperature/12345678", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestTemperatureRecoveryLocation_InvalidCep(t *testing.T) {
	app := fiber.New()
	app.Get("/temperature/:number", handlers.TemperatureRecoveryLocation)

	req := httptest.NewRequest("GET", "/temperature/123", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestTemperatureRecoveryLocation_MissingCep(t *testing.T) {
	app := fiber.New()
	app.Get("/temperature/:number", handlers.TemperatureRecoveryLocation)

	req := httptest.NewRequest("GET", "/temperature/", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
