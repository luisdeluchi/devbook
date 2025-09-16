package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa uma estrutura de erro que será retornada como resposta JSON.
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON envia uma resposta JSON com o código de status especificado e os dados fornecidos.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Erro ao gerar JSON de resposta:", err)
	}
}

// TratarStatusCodeDeErro trata o status code de erro retornado pela API e envia uma resposta JSON com a mensagem de erro.
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI

	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)

}
