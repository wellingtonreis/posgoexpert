package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	config "service_b/internal/config"
)

type ServiceCepImpl struct{}

type Location struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (s ServiceCepImpl) ViaCep(number string) (Location, error) {

	host := config.GetEnv("VIA_CEP_URL", "https://viacep.com.br")
	if host == "" {
		return Location{}, errors.New("invalid host for ViaCEP")
	}

	url := fmt.Sprintf("%s/ws/%s/json/", host, number)

	resp, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("failed to request ViaCEP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, errors.New("invalid response from ViaCEP")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, fmt.Errorf("failed to read response from ViaCEP: %w", err)
	}

	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, errors.New("failed to parse response from ViaCEP")
	}

	return location, nil
}
