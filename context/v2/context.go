package v2

import (
	"context"
	"fmt"
	"net/http"
)

// Pass through the context to our Store and let it be responsible
// This way it can propagate the contexts through to its dependants and they too can be
// responsible for stopping themselves.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// takes in a Store and returns us a http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error
		}

		fmt.Fprint(w, data)
	}
}

/* Before: Does it make sense for our web server to be concerned with manually
* cancelling Store? What if Store also happens to depend on other slow-running
* processes? We'll have to make sure that Store.Cancel correctly propagates the
* cancellation to all of its dependants.
* */
