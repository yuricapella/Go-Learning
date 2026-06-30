package exampleapp

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type TokenizeResponse struct {
	Tokens []string `json:"tokens"`
	Count  int      `json:"count"`
}

func TokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}

	texto, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	tokens := strings.Fields(string(texto))
	resposta := TokenizeResponse{
		Tokens: tokens,
		Count:  len(tokens),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}
