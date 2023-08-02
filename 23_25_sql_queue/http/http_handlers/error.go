package http_handlers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func addError(message string, code int, w http.ResponseWriter) {
	err := Error{message}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
