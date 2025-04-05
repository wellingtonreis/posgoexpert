package handlers_test

import (
	"bytes"
	"errors"
	"net/http/httptest"
	dto "service_a/internal/dto"
	service "service_a/internal/service"
	usecase "service_a/internal/usecase"
	"testing"

	handlers "service_a/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type MockCepUseCase struct {
	mock.Mock
}

func (m *MockCepUseCase) GetCep(number string) (*dto.CepDTO, error) {
	args := m.Called(number)
	return args.Get(0).(*dto.CepDTO), args.Error(1)
}

func TestGetCep(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockCepUseCase) // Inicializando o mock
	cepService := service.NewServiceCepImpl(nil, "http://serviceb:9000")
	cepUseCase := usecase.NewCepUseCase(cepService)
	handler := handlers.NewCepHandler(*cepUseCase)

	app.Post("/cep", handler.GetCep)

	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/cep", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Invalid CEP format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/cep", bytes.NewReader([]byte(`{"number": "123"}`)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Error fetching CEP", func(t *testing.T) {
		mockUseCase.On("GetCep", "12345678").Return(nil, errors.New("service error"))

		req := httptest.NewRequest("POST", "/cep", bytes.NewReader([]byte(`{"number": "12345678"}`)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Successful fetch", func(t *testing.T) {
		expectedResponse := &dto.CepDTO{Number: "12345678"}
		mockUseCase.On("GetCep", "12345678").Return(expectedResponse, nil)

		req := httptest.NewRequest("POST", "/cep", bytes.NewReader([]byte(`{"number": "12345678"}`)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
