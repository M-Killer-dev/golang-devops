package controllers

import (
	"encoding/json"
	"fmt"
	"net"
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

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	ip_address := r.FormValue("ip")

	fmt.Printf("IP address => %s\n", ip_address)

	if net.ParseIP(ip_address) != nil {
		w.Write([]byte("Valid IP address"))
	} else {
		w.Write([]byte("Invalid IP address"))
	}
}
