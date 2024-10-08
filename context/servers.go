package main

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Print(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
