package methods

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Methods() {
	r := chi.NewRouter()

	var countOfRequest int = 0
	r.Route("/hello", func(r chi.Router) {
		r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := chi.URLParam(r, "name")
			countOfRequest++
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"Hello": name})
		})

		r.Get("/count", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"Hello count": strconv.Itoa(countOfRequest)})

		})

		r.Put("/{delete_count:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			numberToDelete := chi.URLParam(r, "delete_count")
			numberToDeleteINT, _ := strconv.Atoi(numberToDelete)
			if numberToDeleteINT > countOfRequest {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			countOfRequest -= numberToDeleteINT
		})
	})

	http.ListenAndServe(":8080", r)
}
