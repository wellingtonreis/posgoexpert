package service_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	service "service_b/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViaCep_Success(t *testing.T) {
	expectedLocation := service.Location{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedLocation)
	}))
	defer server.Close()

	os.Setenv("HOST_VIA_CEP", server.URL)
	defer os.Unsetenv("HOST_VIA_CEP")

	serviceCep := service.ServiceCepImpl{}
	location, err := serviceCep.ViaCep("01001000")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if location != expectedLocation {
		t.Fatalf("expected %v, got %v", expectedLocation, location)
	}
}

func TestViaCep_InvalidHost(t *testing.T) {
	os.Setenv("HOST_VIA_CEP", "https://viacepp.com.br/")
	defer os.Unsetenv("HOST_VIA_CEP")

	serviceCep := service.ServiceCepImpl{}
	_, err := serviceCep.ViaCep("01001000")

	assert.NotNil(t, err)
}

func TestViaCep_InvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("HOST_VIA_CEP", server.URL)
	defer os.Unsetenv("HOST_VIA_CEP")

	serviceCep := service.ServiceCepImpl{}
	_, err := serviceCep.ViaCep("01001000")
	if err == nil || err.Error() != "invalid response from ViaCEP" {
		t.Fatalf("expected 'invalid response from ViaCEP' error, got %v", err)
	}
}

func TestViaCep_FailedToReadResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	os.Setenv("HOST_VIA_CEP", server.URL)
	defer os.Unsetenv("HOST_VIA_CEP")

	serviceCep := service.ServiceCepImpl{}
	_, err := serviceCep.ViaCep("01001000")
	if err == nil || err.Error() != "failed to parse response from ViaCEP" {
		t.Fatalf("expected 'failed to parse response from ViaCEP' error, got %v", err)
	}
}
