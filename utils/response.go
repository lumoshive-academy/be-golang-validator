package utils

import (
	"encoding/json"
	"net/http"
)

type Reponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Reponse2 struct {
	Status  bool `json:"status"`
	Message any  `json:"message"`
	Data    any  `json:"data,omitempty"`
}

func ResponseSuccess(w http.ResponseWriter, code int, message string, data any) {
	response := Reponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func ResponseBadRequest(w http.ResponseWriter, code int, message string) {
	response := Reponse{
		Status:  true,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func ResponseBadRequest2(w http.ResponseWriter, code int, message any) {
	response := Reponse2{
		Status:  true,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
