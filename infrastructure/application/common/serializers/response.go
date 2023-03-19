package serializers

import (
	"encoding/json"
	"net/http"
	"quiz-1/business/core"
)

func BuildResponse(w http.ResponseWriter, message string, data interface{}, status int) {
	response := core.Response{
		Message: message,
		Data:    data,
		Status:  status,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
