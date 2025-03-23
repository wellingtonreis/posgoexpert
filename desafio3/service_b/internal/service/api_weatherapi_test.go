package service_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	service "service_b/internal/service"
	"testing"
)

func TestWeatherApi_Success(t *testing.T) {
	// Mock server
	mockResponse := `{"location": {"region": "TestRegion"}, "current": {"temp_c": 25.0, "temp_f": 77.0}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Set environment variables
	os.Setenv("HOST_WEATHER_API", server.URL)
	os.Setenv("KEY_WEATHER_API", "testkey")

	service := service.ServiceWeatherApiImpl{}
	weather, err := service.WeatherApi("TestLocation")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if weather.Location.Region != "TestRegion" {
		t.Errorf("expected region 'TestRegion', got %v", weather.Location.Region)
	}

	if weather.Current.TempC != 25.0 {
		t.Errorf("expected temp_c 25.0, got %v", weather.Current.TempC)
	}

	if weather.Current.TempF != 77.0 {
		t.Errorf("expected temp_f 77.0, got %v", weather.Current.TempF)
	}
}

func TestWeatherApi_InvalidHost(t *testing.T) {
	os.Setenv("HOST_WEATHER_API", "")
	os.Setenv("KEY_WEATHER_API", "testkey")

	service := service.ServiceWeatherApiImpl{}
	_, err := service.WeatherApi("TestLocation")

	if err == nil || err.Error() != "invalid host for WeatherAPI" {
		t.Fatalf("expected 'invalid host for WeatherAPI' error, got %v", err)
	}
}

func TestWeatherApi_InvalidKey(t *testing.T) {
	os.Setenv("HOST_WEATHER_API", "http://api.weatherapi.com/v1/current.json")
	os.Setenv("KEY_WEATHER_API", "")

	service := service.ServiceWeatherApiImpl{}
	_, err := service.WeatherApi("TestLocation")

	if err == nil || err.Error() != "invalid key for WeatherAPI" {
		t.Fatalf("expected 'invalid key for WeatherAPI' error, got %v", err)
	}
}

func TestWeatherApi_FailedRequest(t *testing.T) {
	os.Setenv("HOST_WEATHER_API", "http://invalidhost")
	os.Setenv("KEY_WEATHER_API", "testkey")

	service := service.ServiceWeatherApiImpl{}
	_, err := service.WeatherApi("TestLocation")

	if err == nil || err.Error() != "failed to request WeatherAPI" {
		t.Fatalf("expected 'failed to request WeatherAPI' error, got %v", err)
	}
}

func TestWeatherApi_InvalidResponse(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Set environment variables
	os.Setenv("HOST_WEATHER_API", server.URL)
	os.Setenv("KEY_WEATHER_API", "testkey")

	service := service.ServiceWeatherApiImpl{}
	_, err := service.WeatherApi("TestLocation")

	if err == nil || err.Error() != "invalid response from WeatherAPI" {
		t.Fatalf("expected 'invalid response from WeatherAPI' error, got %v", err)
	}
}
