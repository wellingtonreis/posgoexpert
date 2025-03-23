package usecase_test

import (
	"errors"
	domain "service_b/internal/domain"
	service "service_b/internal/service"
	usecase "service_b/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCepService struct {
	mock.Mock
}

func (m *MockCepService) ViaCep(number string) (service.Location, error) {
	args := m.Called(number)
	return args.Get(0).(service.Location), args.Error(1)
}

type MockWeatherApi struct {
	mock.Mock
}

func (m *MockWeatherApi) WeatherApi(q string) (service.Weather, error) {
	args := m.Called(q)
	return args.Get(0).(service.Weather), args.Error(1)
}

func TestGetTemperatureRecoveryLocation(t *testing.T) {
	mockCepService := new(MockCepService)
	mockWeatherApi := new(MockWeatherApi)
	useCase := usecase.NewTemperatureRecoveryLocationUseCase(mockCepService, mockWeatherApi)

	cep := "01001000"
	cidade := "São Paulo"
	location := service.Location{Cep: cep, Localidade: cidade}
	temperatureC := 25.0
	temperatureF := temperatureC*1.8 + 32
	weather := service.Weather{Current: struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	}{TempC: temperatureC, TempF: temperatureF}}

	mockCepService.On("ViaCep", cep).Return(location, nil)
	mockWeatherApi.On("WeatherApi", cidade).Return(weather, nil)

	expected := domain.TemperatureLocation{
		Number:     cep,
		City:       cidade,
		Celsius:    float32(temperatureC),
		Fahrenheit: float32(temperatureF),
		Kelvin:     float32(temperatureC + 273.15),
	}

	result, err := useCase.GetTemperatureRecoveryLocation(cep)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetTemperatureRecoveryLocation_ErrorCepService(t *testing.T) {
	mockCepService := new(MockCepService)
	mockWeatherApi := new(MockWeatherApi)
	useCase := usecase.NewTemperatureRecoveryLocationUseCase(mockCepService, mockWeatherApi)

	cep := "01001000"
	expectedErr := errors.New("CEP not found")
	mockCepService.On("ViaCep", cep).Return(service.Location{}, expectedErr)

	_, err := useCase.GetTemperatureRecoveryLocation(cep)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestGetTemperatureRecoveryLocation_ErrorWeatherApi(t *testing.T) {
	mockCepService := new(MockCepService)
	mockWeatherApi := new(MockWeatherApi)
	useCase := usecase.NewTemperatureRecoveryLocationUseCase(mockCepService, mockWeatherApi)

	cep := "01001000"
	cidade := "São Paulo"
	location := service.Location{Cep: cep, Localidade: cidade}
	expectedErr := errors.New("Weather API unavailable")

	mockCepService.On("ViaCep", cep).Return(location, nil)
	mockWeatherApi.On("WeatherApi", cidade).Return(service.Weather{}, expectedErr)

	_, err := useCase.GetTemperatureRecoveryLocation(cep)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}
