package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem vindo a minha pagina!")
	})
	http.ListenAndServe(":8080", nil)
}
