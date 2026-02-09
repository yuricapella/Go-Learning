package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(responseWriter http.ResponseWriter, statusCode int, dados interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(responseWriter).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Erro retorna um erro em formato JSON para a requisição
func Erro(responseWriter http.ResponseWriter, statusCode int, erro error) {
	JSON(responseWriter, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})

}
