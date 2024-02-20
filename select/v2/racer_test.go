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

	t.Run("compares speeds of servers, returnin url of the fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

	// defer prefix will call a function at the end of a containing function
	// Keep instruction near where you created server for easier spotting
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL //"http://www.facebook.com"
		fastURL := fastServer.URL //before:"http://www.quii.dev"

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respons within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

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
