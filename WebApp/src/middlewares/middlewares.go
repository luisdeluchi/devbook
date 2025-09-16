package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger é um middleware que registra as requisições recebidas
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logica de log
		log.Printf("Requisição recebida: %s %s %s", r.Method, r.RequestURI, r.Host)

		// Chama a próxima função no middleware
		proximaFuncao(w, r)
	}
}

// Autenticar verifica existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Chama a próxima função no middleware
		proximaFuncao(w, r)
	}
}
