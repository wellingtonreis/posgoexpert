package dto

type CepDTO struct {
	Number string `json:"cep" validate:"required"`
}
