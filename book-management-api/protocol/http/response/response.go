package response

import (
	"book-management-api/domain/dto"
	"encoding/json"
	"net/http"
)

// Utility functions
func SendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	SendJSONResponse(w, dto.ErrorResponse{
		Status: "error",
		Error:  message,
	}, statusCode)
}
