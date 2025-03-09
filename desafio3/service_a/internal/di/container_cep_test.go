package di_test

import (
	di "service_a/internal/di"
	"testing"

	handlers "service_a/internal/handlers"
	service "service_a/internal/service"
	usecase "service_a/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestBuildContainerCep(t *testing.T) {
	container, err := di.BuildContainerCep()
	assert.NoError(t, err)
	assert.NotNil(t, container)
	assert.NotNil(t, container.CepUseCase)
	assert.NotNil(t, container.CepHandler)

	expectedService := service.ServiceCepImpl{
		HTTPClient: nil,
		BaseURL:    "http://localhost:9000",
	}
	expectedUseCase := usecase.NewCepUseCase(expectedService)
	assert.Equal(t, expectedUseCase, container.CepUseCase)

	expectedHandler := handlers.NewCepHandler(*expectedUseCase)
	assert.Equal(t, expectedHandler, container.CepHandler)
}
