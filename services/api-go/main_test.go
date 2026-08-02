package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthEndpoint(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rr := httptest.NewRecorder()

    healthHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rr.Code)
    }

    var payload map[string]string
    if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
        t.Fatalf("expected valid JSON payload: %v", err)
    }

    if payload["status"] != "ok" || payload["service"] != "api-go" {
        t.Fatalf("unexpected payload: %#v", payload)
    }
}

func TestCommentsEndpoint(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/comments", nil)
    rr := httptest.NewRecorder()

    commentsHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rr.Code)
    }

    var payload []Comment
    if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
        t.Fatalf("expected valid JSON payload: %v", err)
    }

    if len(payload) < 2 {
        t.Fatalf("expected at least 2 comments, got %d", len(payload))
    }
}
