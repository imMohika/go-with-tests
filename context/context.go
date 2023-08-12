package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())

		if err != nil {
			println(err.Error())
			return
		}
		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
