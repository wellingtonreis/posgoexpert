package domain

type TemperatureLocation struct {
	Number     string  `json:"cep"`
	City       string  `json:"city"`
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
	Kelvin     float32 `json:"kelvin"`
}
