package exampleapp

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const jsonContentType = "application/json"

type WeatherRecord struct {
	City        string  `json:"city"`
	TempCelsius float64 `json:"temp_celsius"`
}

type AddResponse struct {
	OK         bool `json:"ok"`
	NumRecords int  `json:"num_records"`
}

type WeatherStore struct {
	mu      sync.Mutex
	records []WeatherRecord
}

func (s *WeatherStore) Add(record WeatherRecord) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.records = append(s.records, record)
	return len(s.records)
}

func (s *WeatherStore) AddHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var record WeatherRecord
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		log.Printf("unmarshal: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total := s.Add(record)
	response := AddResponse{
		OK:         true,
		NumRecords: total,
	}

	w.Header().Set("Content-Type", jsonContentType)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("marshal: %s", err)
	}
}
