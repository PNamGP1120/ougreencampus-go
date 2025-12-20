package common

import "net/http"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func Message(msg string) Response {
	return Response{
		Success: true,
		Message: msg,
	}
}

func Error(message string, err interface{}) (int, Response) {
	return http.StatusBadRequest, Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}
