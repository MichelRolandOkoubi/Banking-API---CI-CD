// src/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	account := Account{ID: "ACC001", Balance: 15000.50}
	json.NewEncoder(w).Encode(account)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/api/balance", getBalance)
	http.HandleFunc("/health", healthCheck)
	log.Println("Banking API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
