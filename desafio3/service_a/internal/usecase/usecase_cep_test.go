package usecase_test

import (
	"errors"
	domain "service_a/internal/domain"
	usecase "service_a/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServiceCep struct {
	mock.Mock
}

func (m *MockServiceCep) ServiceCep(number string) (domain.TemperatureLocation, error) {
	args := m.Called(number)
	return args.Get(0).(domain.TemperatureLocation), args.Error(1)
}

func TestGetCep(t *testing.T) {
	t.Run("should return temperature location when service succeeds", func(t *testing.T) {
		mockSvc := new(MockServiceCep)
		expectedResult := domain.TemperatureLocation{
			Number:     "01001000",
			City:       "São Paulo",
			Celsius:    25.5,
			Fahrenheit: 77.9,
			Kelvin:     298.6,
		}
		mockSvc.On("ServiceCep", "01001000").Return(expectedResult, nil)

		useCase := usecase.NewCepUseCase(mockSvc)

		result, err := useCase.GetCep("01001000")

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		mockSvc.AssertExpectations(t)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		mockSvc := new(MockServiceCep)
		expectedError := errors.New("service error")
		mockSvc.On("ServiceCep", "invalid").Return(domain.TemperatureLocation{}, expectedError)

		useCase := usecase.NewCepUseCase(mockSvc)
		result, err := useCase.GetCep("invalid")

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Empty(t, result)
		mockSvc.AssertExpectations(t)
	})
}

func TestNewCepUseCase(t *testing.T) {
	mockSvc := new(MockServiceCep)
	useCase := usecase.NewCepUseCase(mockSvc)
	assert.NotNil(t, useCase)
}
