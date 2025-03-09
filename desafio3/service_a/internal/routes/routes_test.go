package routes_test

import (
	"net/http/httptest"
	routes "service_a/internal/routes"
	"strings"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	assert "github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()

	routes.SetupRoutes(app)

	req := httptest.NewRequest("POST", "/api/v1/cep/post", strings.NewReader(`{"cep":"01001000"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
