package dto

import (
	"testing"

	validator "github.com/go-playground/validator/v10"
)

func TestCepDTO_Validation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		cep     CepDTO
		wantErr bool
	}{
		{
			name:    "valid cep",
			cep:     CepDTO{Number: "12345678"},
			wantErr: false,
		},
		{
			name:    "empty cep",
			cep:     CepDTO{Number: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("CepDTO validation error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
