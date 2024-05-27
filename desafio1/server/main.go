package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fastjson"
)

type Currency struct {
	Code       string    `json:"code"`
	Codein     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varbid"`
	PctChange  string    `json:"pctchange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"createdate"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/cotacao", CurrencyHandler)

	db, err := sql.Open("sqlite3", "./currency.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS currency ( bid VARCHAR(255) );")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(100 * time.Millisecond):

		res, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
		if err != nil {
			http.Error(w, "Não foi possível processar a requisição", http.StatusInternalServerError)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Erro ao ler o corpo da resposta:", err)
			return
		}

		var p fastjson.Parser
		value, err := p.ParseBytes(body)
		if err != nil {
			http.Error(w, "Erro ao tentar converter os dados em json", http.StatusInternalServerError)
			return
		}

		currencyJSON := value.Get("USDBRL").String()
		var currency Currency
		if err := json.Unmarshal([]byte(currencyJSON), &currency); err != nil {
			panic(err)
		}

		jsonOutput := fmt.Sprintf(`{"bid": "%s"}`, currency.Bid)

		fmt.Println(jsonOutput)

		db, err := sql.Open("sqlite3", "./currency.db")
		if err != nil {
			log.Println("Erro ao tentar conectar no banco de dados:", err)
			http.Error(w, "Erro ao tentar conectar no banco de dados!", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer dbCancel()

		_, err = db.ExecContext(dbCtx, "INSERT INTO currency (bid) VALUES (?)", currency.Bid)
		if err != nil {
			log.Println("Erro ao tentar inserir um novo registro no banco de dados:", err)
			http.Error(w, "Erro ao tentar inserir um novo registro no banco de dados!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonOutput))

	case <-ctx.Done():
		http.Error(w, "A requisição foi cancelada!", http.StatusInternalServerError)
	}
}
