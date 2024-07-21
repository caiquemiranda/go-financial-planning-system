package http

import (
	"net/http"

	"github.com/caiquemiranda/go-financial-planning-system/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)
	http.ListenAndServe(":8080", nil)
}
