package main

import (
	"fmt"
	"time"
)

// EXECUTE COM SEGUINTE COMANDO "go run main.go brasilapi.go viacep.go"
func main() {
	cep := "01153000"
	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	canalBrasilApi := make(chan RespostaBrasilApi)
	canalViaCep := make(chan RespostaViaCep)

	go BuscarCepBrasilApi(urlBrasilApi, canalBrasilApi)
	go BuscarCepViaCep(urlViaCep, canalViaCep)

	select {
	case respostaBrasilApi := <-canalBrasilApi:
		fmt.Printf("Brasil API: %+v\n", respostaBrasilApi)
	case respostaViaCep := <-canalViaCep:
		fmt.Printf("Via Cep: %+v\n", respostaViaCep)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout: O tempo expirou!")
	}
}
