package httpserver

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(encoded)
	return err
}

func WriteError(w http.ResponseWriter, status int, code, message string) {
	response := errorResponse{
		Error: errorBody{
			Code:    code,
			Message: message,
		},
	}
	_ = WriteJSON(w, status, response)
}
