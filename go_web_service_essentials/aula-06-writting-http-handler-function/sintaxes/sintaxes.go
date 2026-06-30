package sintaxes

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func NovoRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	return mux
}

func IniciarServidor(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
