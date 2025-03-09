package domain

import (
	"encoding/json"
	"testing"
)

func TestTemperatureLocationJSON(t *testing.T) {
	tl := TemperatureLocation{
		Number:     "12345",
		City:       "Test City",
		Celsius:    25.0,
		Fahrenheit: 77.0,
		Kelvin:     298.15,
	}

	data, err := json.Marshal(tl)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	var unmarshalledTL TemperatureLocation
	err = json.Unmarshal(data, &unmarshalledTL)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if tl != unmarshalledTL {
		t.Errorf("Expected %v, got %v", tl, unmarshalledTL)
	}
}

func TestTemperatureLocationFields(t *testing.T) {
	tl := TemperatureLocation{
		Number:     "12345",
		City:       "Test City",
		Celsius:    25.0,
		Fahrenheit: 77.0,
		Kelvin:     298.15,
	}

	if tl.Number != "12345" {
		t.Errorf("Expected Number to be '12345', got %s", tl.Number)
	}
	if tl.City != "Test City" {
		t.Errorf("Expected City to be 'Test City', got %s", tl.City)
	}
	if tl.Celsius != 25.0 {
		t.Errorf("Expected Celsius to be 25.0, got %f", tl.Celsius)
	}
	if tl.Fahrenheit != 77.0 {
		t.Errorf("Expected Fahrenheit to be 77.0, got %f", tl.Fahrenheit)
	}
	if tl.Kelvin != 298.15 {
		t.Errorf("Expected Kelvin to be 298.15, got %f", tl.Kelvin)
	}
}
