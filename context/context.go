package contexts

import (
	"fmt"
	"net/http"
)

// store fetches data and cancels
type Store interface {
	Fetch() string
	Cancel()
}

// takes in a Store and returns us a http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// result store.Fetch() is written to http response
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
