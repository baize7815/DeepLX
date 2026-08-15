package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	router := Router(&Config{})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got, want := recorder.Body.String(), "{\"status\":\"ok\"}"; got != want {
		t.Fatalf("expected body %q, got %q", want, got)
	}
}

func TestHealthzHead(t *testing.T) {
	router := Router(&Config{})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodHead, "/healthz", nil)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if recorder.Body.Len() != 0 {
		t.Fatalf("expected an empty body, got %q", recorder.Body.String())
	}
}
