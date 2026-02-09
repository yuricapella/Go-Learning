package util

import (
	"encoding/json"
	"net/http"
)

// ParseJSONRequestBody decodifica o corpo da requisição HTTP para o tipo esperado.
func ParseJSONRequestBody(request *http.Request, target interface{}) error {
	return json.NewDecoder(request.Body).Decode(target)
}

// WriteJSONResponse serializa o payload como JSON, define o status desejado, e envia ao cliente.
func WriteJSONResponse(responseWriter http.ResponseWriter, statusCode int, responsePayload interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)

	if encodingError := json.NewEncoder(responseWriter).Encode(responsePayload); encodingError != nil {
		http.Error(responseWriter, "Erro ao serializar resposta em JSON: "+encodingError.Error(), http.StatusInternalServerError)
	}
}
