package controllers

import (
	"encoding/json"
	"net/http"
	"time"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Version    string `json:"version"`
		Date       string `json:"date"`
		Kubernetes bool   `json:"kubernetes"`
	}

	response := Response{
		Version:    "0.1.0",
		Date:       time.Now().Format(time.RFC3339),
		Kubernetes: false,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
