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

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// defer prefix will call a function at the end of a containing function
	// Keep instruction near where you created server for easier spotting
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL //"http://www.facebook.com"
	fastURL := fastServer.URL //before:"http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}



// Handler func takes in a response writer & req
// This is how we write a server in Go
// Difference is we are wrapping it in httptest.NewServer, makes it easier to use in tests
// It finds an open port to listen to & we can close servers when done with tests
func makeDelayedServer(delay time.Duration) *httptest.Server {
        return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                time.Sleep(delay)
                w.WriteHeader(http.StatusOK)
        }))
}
