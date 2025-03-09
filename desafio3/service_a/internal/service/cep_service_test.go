package service_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "service_a/internal/domain"
	service "service_a/internal/service"

	"github.com/stretchr/testify/assert"
)

// Mock do HTTPClient
func mockServer(responseBody string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	})
	return httptest.NewServer(handler)
}

func TestServiceCep(t *testing.T) {
	t.Run("should return temperature location when service succeeds", func(t *testing.T) {
		mock := mockServer(`{"Number":"","City":"","Celsius":0,"Fahrenheit":0,"Kelvin":0}`, http.StatusOK)
		defer mock.Close()

		serviceCep := service.ServiceCepImpl{
			HTTPClient: mock.Client(),
			BaseURL:    mock.URL,
		}

		expectedResult := domain.TemperatureLocation{
			Number:     "", // 01001-000
			City:       "", // São Paulo
			Celsius:    0,  // 29.1
			Fahrenheit: 0,  // 84.38
			Kelvin:     0,  // 302.25
		}

		result, err := serviceCep.ServiceCep("01001000")
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return error when response is invalid JSON", func(t *testing.T) {
		mock := mockServer(`{invalid json}`, http.StatusOK)
		defer mock.Close()

		serviceCep := service.ServiceCepImpl{
			HTTPClient: mock.Client(),
			BaseURL:    mock.URL,
		}

		result, err := serviceCep.ServiceCep("01001000")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse temperature location")
		assert.Empty(t, result)
	})

	t.Run("should return error when service returns non-200 status code", func(t *testing.T) {
		mock := mockServer(``, http.StatusInternalServerError)
		defer mock.Close()

		serviceCep := service.ServiceCepImpl{
			HTTPClient: mock.Client(),
			BaseURL:    mock.URL,
		}

		result, err := serviceCep.ServiceCep("01001000")
		assert.Error(t, err)
		assert.Equal(t, "invalid response", err.Error())
		assert.Empty(t, result)
	})

	t.Run("should return error when service fails to get response", func(t *testing.T) {
		brokenClient := &http.Client{
			Transport: &errorTransport{},
		}

		serviceCep := service.ServiceCepImpl{
			HTTPClient: brokenClient,
			BaseURL:    "http://localhost:9000/",
		}

		result, err := serviceCep.ServiceCep("01001000")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get temperature location")
		assert.Empty(t, result)
	})

	t.Run("should return error when response body cannot be read", func(t *testing.T) {
		unreadableClient := &http.Client{
			Transport: &unreadableBodyTransport{},
		}

		serviceCep := service.ServiceCepImpl{
			HTTPClient: unreadableClient,
			BaseURL:    "http://localhost:9000/",
		}

		result, err := serviceCep.ServiceCep("01001000")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to read response body")
		assert.Empty(t, result)
	})
}

type errorTransport struct{}

func (e *errorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("network error")
}

type unreadableBodyTransport struct{}

func (u *unreadableBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(&errorReader{}), // Corpo que falha ao ser lido
	}, nil
}

type errorReader struct{}

func (errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("failed to read response body")
}

func (errorReader) Close() error {
	return nil
}
