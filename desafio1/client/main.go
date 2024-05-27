package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CurrencyUsd struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao realizar a requisição!", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("A requisição falhou!", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Falha ao receber a resposta", err)
	}

	log.Printf("Resposta: %s", body)
	jsonData := []byte(body)

	var currency CurrencyUsd
	if err := json.Unmarshal(jsonData, &currency); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", currency.Bid))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Cotação salva em cotacao.txt com sucesso!")
}
