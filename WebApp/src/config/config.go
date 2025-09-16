package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ApiUrl representa a URL para comunicação com a API.
	APIURL = ""
	// Porta representa a porta do servidor.
	Porta = 0
	// HashKey representa a chave de hash para segurança.
	HashKey []byte
	// BlockKey representa a chave de bloqueio para segurança.
	BlockKey []byte
)

// Carregar inicializa as configurações do aplicativo.
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)

	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
