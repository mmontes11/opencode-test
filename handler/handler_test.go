package handler

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheck(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()

    HealthCheck(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200, got %d", resp.StatusCode)
    }

    ct := resp.Header.Get("Content-Type")
    if ct != "application/json" {
        t.Fatalf("expected Content-Type application/json, got %s", ct)
    }

    var body map[string]string
    if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
        t.Fatalf("failed to decode response body: %v", err)
    }

    if status, ok := body["status"]; !ok || status != "ok" {
        t.Fatalf("expected body {\"status\": \"ok\"}, got %v", body)
    }
}
