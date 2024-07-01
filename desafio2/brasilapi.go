package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespostaBrasilApi struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func BuscarCepBrasilApi(url string, ch chan<- RespostaBrasilApi) {
	resposta, erro := http.Get(url)
	if erro != nil {
		fmt.Printf("Erro na requisição: %v\n", erro)
		close(ch)
		return
	}
	defer resposta.Body.Close()

	resultado, erro := io.ReadAll(resposta.Body)
	if erro != nil {
		fmt.Printf("Erro no resultado da busca: %v\n", erro)
		close(ch)
		return
	}

	var cepBrasilApi RespostaBrasilApi
	erro = json.Unmarshal(resultado, &cepBrasilApi)
	if erro != nil {
		fmt.Printf("Erro no parser do json: %v\n", erro)
		close(ch)
		return
	}

	ch <- cepBrasilApi
	close(ch)
}
