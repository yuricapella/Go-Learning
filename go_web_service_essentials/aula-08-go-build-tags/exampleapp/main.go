package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/feature-flags", featureFlagsHandler)
	mux.HandleFunc("/by", byHandler)

	log.Println("servidor em http://localhost:8080")
	log.Println("rotas: /health, /feature-flags, /by")
	if profileEnabled() {
		log.Println("rota extra: /debug/pprof/")
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func featureFlagsHandler(w http.ResponseWriter, _ *http.Request) {
	resposta := map[string]any{
		"profile_enabled": profileEnabled(),
		"goos_founder":    founder(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func byHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(founder()))
}
