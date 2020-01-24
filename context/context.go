package main

import (
	"context"
	"fmt"
	"net/http"
)

// Store interface for our store
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server is a handler that fetches the store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprintf(w, data)
	}
}
