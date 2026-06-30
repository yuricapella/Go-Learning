package exampleapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestTokenizeHandler(t *testing.T) {
	body := strings.NewReader("Have you seen Who's on First?")
	request := httptest.NewRequest(http.MethodPost, "/tokenize", body)
	recorder := httptest.NewRecorder()

	TokenizeHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status code = %d, want %d", response.StatusCode, http.StatusOK)
	}

	var got TokenizeResponse
	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	want := TokenizeResponse{
		Tokens: []string{"Have", "you", "seen", "Who's", "on", "First?"},
		Count:  6,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("response = %#v, want %#v", got, want)
	}
}

func TestTokenizeHandlerRejectsGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/tokenize", nil)
	recorder := httptest.NewRecorder()

	TokenizeHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("status code = %d, want %d", response.StatusCode, http.StatusMethodNotAllowed)
	}
}
