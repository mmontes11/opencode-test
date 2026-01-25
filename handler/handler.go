package handler

import (
    "encoding/json"
    "net/http"
)

// HealthCheck returns a simple JSON response to indicate the service is running.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
