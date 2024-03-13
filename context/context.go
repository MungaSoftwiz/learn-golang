package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}


// takes in a Store and returns us a http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
