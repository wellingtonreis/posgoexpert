package di

import (
	config "service_a/internal/config"
	handlers "service_a/internal/handlers"
	service "service_a/internal/service"
	usecase "service_a/internal/usecase"
)

type ContainerCep struct {
	CepUseCase *usecase.CepUseCase
	CepHandler *handlers.CepHandler
}

func BuildContainerCep() (*ContainerCep, error) {

	environment := config.GetEnv("ENVIRONMENT", "production")

	url := "http://localhost:9000"
	if environment == "development" {
		url = "http://serviceb:9000"
	}

	cepService := service.NewServiceCepImpl(nil, url)
	cepUseCase := usecase.NewCepUseCase(cepService)
	cepHandler := handlers.NewCepHandler(*cepUseCase)

	return &ContainerCep{
		CepUseCase: cepUseCase,
		CepHandler: cepHandler,
	}, nil
}
