package context

import (
	"fmt"
	"log"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprint(w, data)
	}
}
