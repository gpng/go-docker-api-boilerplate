package utils

import (
	"encoding/json"
	"net/http"
)

// Message maps and status and message into JSON formatted string
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond sends a JSON response to a http request
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
