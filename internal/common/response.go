package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func Success(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func Error(message string, err interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}
