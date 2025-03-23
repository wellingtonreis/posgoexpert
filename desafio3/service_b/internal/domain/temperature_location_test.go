package domain_test

import (
	"testing"

	domain "service_b/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestTemperatureLocation(t *testing.T) {
	t.Run("Test TemperatureLocation struct", func(t *testing.T) {
		tempLoc := domain.TemperatureLocation{
			Number:     "12345",
			City:       "Test City",
			Celsius:    25.0,
			Fahrenheit: 77.0,
			Kelvin:     298.15,
		}

		assert.Equal(t, "12345", tempLoc.Number)
		assert.Equal(t, "Test City", tempLoc.City)
		assert.Equal(t, float32(25.0), tempLoc.Celsius)
		assert.Equal(t, float32(77.0), tempLoc.Fahrenheit)
		assert.Equal(t, float32(298.15), tempLoc.Kelvin)
	})
}
