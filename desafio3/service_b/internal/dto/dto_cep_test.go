package dto_test

import (
	"testing"

	dto "service_b/internal/dto"

	validator "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCepDTO_Valid(t *testing.T) {
	validate := validator.New()
	cep := dto.CepDTO{Number: "12345-678"}

	err := validate.Struct(cep)

	assert.NoError(t, err)
}

func TestCepDTO_Invalid(t *testing.T) {
	validate := validator.New()
	cep := dto.CepDTO{Number: ""}

	err := validate.Struct(cep)

	assert.Error(t, err)
}
