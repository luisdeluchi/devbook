package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	//stringConexaoBanco
	StringConexaoBanco = ""
	// Porta onde a API vai rodar
	Porta = 0

	// ScretKey é a chave secreta usada para assinar os tokens JWT
	ScretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000 // Define a porta padrão se não for possível ler do ambiente
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	ScretKey = []byte(os.Getenv("SECRET_KEY"))

}
