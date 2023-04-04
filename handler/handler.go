package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hello, world!"}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
