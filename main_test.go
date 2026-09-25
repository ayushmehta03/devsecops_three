
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHealthEndpoint(t *testing.T) {
    req := httptest.NewRequest(
        http.MethodGet, "/health", nil,
    )

    recorder := httptest.NewRecorder()
    newRouter().ServeHTTP(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Fatalf("Expected 200, got %d", recorder.Code)
    }

    if !strings.Contains(recorder.Body.String(), "healthy") {
        t.Fatal("Health check response is incorrect")
    }
}
