package sintaxes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

func NovaRequestDeTeste(method, path, body string) *http.Request {
	return httptest.NewRequest(method, path, bytes.NewBufferString(body))
}

func NovoResponseRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func ExecutarHandler(handler http.HandlerFunc, request *http.Request) *http.Response {
	recorder := httptest.NewRecorder()
	handler(recorder, request)
	return recorder.Result()
}

func DecodificarJSON[T any](body io.Reader) (T, error) {
	var valor T
	err := json.NewDecoder(body).Decode(&valor)
	return valor, err
}
