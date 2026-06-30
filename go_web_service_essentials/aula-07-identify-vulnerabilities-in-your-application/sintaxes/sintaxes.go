package sintaxes

import (
	"html/template"
	"io"
	"net/http"
	"strings"
)

func SanitizarEntradaBasica(valor string) string {
	return strings.TrimSpace(valor)
}

func LimitarBody(r *http.Request, maxBytes int64) io.Reader {
	return http.MaxBytesReader(nil, r.Body, maxBytes)
}

func ResponderErroGenerico(w http.ResponseWriter) {
	http.Error(w, "invalid request", http.StatusBadRequest)
}

func EscaparHTML(valor string) string {
	return template.HTMLEscapeString(valor)
}

func MontarQueryParametrizada() (string, []any) {
	query := "SELECT id, email FROM users WHERE email = ?"
	args := []any{"user@example.com"}
	return query, args
}
