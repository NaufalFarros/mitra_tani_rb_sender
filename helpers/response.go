package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	customResponse := Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(customResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops, something went wrong"))
		return
	}

}
