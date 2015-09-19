package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	InternalServer = APIError{
		ID:         "internal_server_error",
		Message:    "Our apologies, but something has gone wrong internally. Please contact support if this problem persists.",
		StatusCode: 500,
	}
)

type APIError struct {
	ID         string `json:"id"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%v %v: %v", e.StatusCode, e.ID, e.Message)
}

func (e APIError) WriteHTTP(w http.ResponseWriter) {
	encoded, err := json.Marshal(e)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(e.StatusCode)
		w.Write(encoded)
	}
}
