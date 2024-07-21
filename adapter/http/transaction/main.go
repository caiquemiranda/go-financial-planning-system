package transaction

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	transaction "github.com/caiquemiranda/go-financial-planning-system/model/transactions"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2024-01-02T15:04:05")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    4152.70,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}
	_ = json.NewEncoder(w).Encode(transactions)

}

func CreateTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

}
