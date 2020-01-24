package main

import (
	"fmt"
	"net/http"
)

// Store interface for our store
type Store interface {
	Fetch() string
	Cancel()
}

// Server is a handler that fetches the store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
