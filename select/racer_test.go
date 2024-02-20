package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// We don't want to depend on external dependency for our tests
// So we use the "net/http/httptest" which enables us create a mock HTTP server
func TestRacer(t *testing.T) {

	// Handler func takes in a response writer & req
	// This is how we write a server in Go
	// Difference is we are wrapping it in httptest.NewServer, makes it easier to use in tests
	// It finds an open port to listen to & we can close servers when done with tests
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL //"http://www.facebook.com"
	fastURL := fastServer.URL //before:"http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
