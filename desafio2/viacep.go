package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespostaViaCep struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func BuscarCepViaCep(url string, ch chan<- RespostaViaCep) {
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

	var cepViaCep RespostaViaCep
	erro = json.Unmarshal(resultado, &cepViaCep)
	if erro != nil {
		fmt.Printf("Erro no parser do json: %v\n", erro)
		close(ch)
		return
	}

	ch <- cepViaCep
	close(ch)
}
