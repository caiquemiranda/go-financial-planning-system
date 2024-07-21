package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github/caiquemiranda/go-financial-planning-system\model\transactions"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createTransactions)	
	http.ListenAndServe(":8080", nil)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2024-01-02T15:04:05")

	var Transactions = Transaction.Transactions{
		Transaction{
			Title:     "Salary",
			Amount:    4152.70,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}
	json.NewEncoder(w).Encode(Transactions)

}

func createTransactions(w http.ResponseWriter, r *http.Request) {

    if r.Method!= "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	w.Header().Set("Content-Type", "application/json")

	var res Transactions{}
    var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)




}
