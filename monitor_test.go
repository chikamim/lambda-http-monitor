package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	err := checkStatus(ts.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestCheckStatusNG(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer ts.Close()

	err := checkStatus(ts.URL)
	if err != ErrResponse {
		t.Fatalf("Should have a response error")
	}
}
