package exampleapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeatherStoreAddHandler(t *testing.T) {
	store := &WeatherStore{}
	body := strings.NewReader(`{"city":"Sao Paulo","temp_celsius":23.5}`)
	request := httptest.NewRequest(http.MethodPost, "/weather", body)
	recorder := httptest.NewRecorder()

	store.AddHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status code = %d, want %d", response.StatusCode, http.StatusOK)
	}

	var got AddResponse
	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	want := AddResponse{OK: true, NumRecords: 1}
	if got != want {
		t.Fatalf("response = %#v, want %#v", got, want)
	}
}

func TestWeatherStoreAddHandlerRejectsInvalidJSON(t *testing.T) {
	store := &WeatherStore{}
	request := httptest.NewRequest(http.MethodPost, "/weather", strings.NewReader(`{invalid`))
	recorder := httptest.NewRecorder()

	store.AddHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
}
