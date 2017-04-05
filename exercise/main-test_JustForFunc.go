package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHandler(t *testing.T) {
    req, err := http.NewRequest(
        http.MethodGet,
        "http://localhost:8080/riceboy",
        nil
    )
    if err != {
        t.Fatalf("could not create request: %v", err)
    }

    rec := httptest.NewRecorder()
    handler(rec, req)

    if rec.Code != http.StatusOK {
        t.Errof("expected status 200; got %d", rec.Code)
    }
    if !strings.Contains(rec.Body.String(), "gopher ricevoy") {
        t.Errorf("unexpected body in response: %q", rec.Body.String())
    }
}
