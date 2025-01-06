package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println(("Request terminada"))
	select {
	case <-time.After(5 * time.Second):
		// Imprime no console stdout
		log.Println("Request Processada com sucesso")

		// Imprime no browser
		w.Write([]byte("Request Processada com Sucesso"))

	case <-ctx.Done():
		log.Println("Request Cancelada")
	}
}
