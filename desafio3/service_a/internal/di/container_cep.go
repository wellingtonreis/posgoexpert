package di

import (
	handlers "service_a/internal/handlers"
	service "service_a/internal/service"
	usecase "service_a/internal/usecase"
)

type ContainerCep struct {
	CepUseCase *usecase.CepUseCase
	CepHandler *handlers.CepHandler
}

func BuildContainerCep() (*ContainerCep, error) {
	cepService := service.ServiceCepImpl{}
	cepUseCase := usecase.NewCepUseCase(cepService)
	cepHandler := handlers.NewCepHandler(cepUseCase)

	return &ContainerCep{
		CepUseCase: cepUseCase,
		CepHandler: &cepHandler,
	}, nil
}
