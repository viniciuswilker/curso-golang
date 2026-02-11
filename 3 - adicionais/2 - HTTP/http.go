package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}
func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func main() {

	// HTTP É O PROTOCOLO DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO WEB

	// CLIENTE FAZ REQUISIÇÃO - SERVIDOR ENVIA RESPOSTA
	// Request - Response

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
