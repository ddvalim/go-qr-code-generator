package response

import (
	"encoding/json"
	"github.com/ddvalim/go-qr-code-generator/core/ports"
	"log"
	"net/http"
)

func Writer(w http.ResponseWriter, statusCode int, contentType string, data interface{}) {
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	Writer(w, statusCode, "application/json", ports.Response{
		Message:    message,
		StatusCode: statusCode,
	})
}
