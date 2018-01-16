package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCheckStatusOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	err := checkStatus(ts.URL, 1*time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestCheckStatusNG(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer ts.Close()

	err := checkStatus(ts.URL, 1*time.Second)
	if err != ErrResponse {
		t.Fatalf("Should have a response error")
	}
}

func TestCheckStatusTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(200)
	}))
	defer ts.Close()

	err := checkStatus(ts.URL, 10*time.Millisecond)
	if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
		t.Fatalf("Should have a timeout error - %v", err)
	}
}
