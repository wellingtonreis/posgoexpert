package usecase

import (
	domain "service_a/internal/domain"
)

type ServiceCep interface {
	ServiceCep(number string) (domain.TemperatureLocation, error)
}

type CepUseCase struct {
	srv ServiceCep
}

func NewCepUseCase(srv ServiceCep) *CepUseCase {
	return &CepUseCase{srv: srv}
}

func (u *CepUseCase) GetCep(number string) (domain.TemperatureLocation, error) {
	return u.srv.ServiceCep(number)
}
