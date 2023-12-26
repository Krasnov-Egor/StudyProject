package methods

//import (
//	"fmt"
//	"github.com/google/uuid"
//)

//
//import (
//	"encoding/json"
//	"github.com/go-chi/chi/v5"
//	"net/http"
//)
//
//func MethodGetCount() {
//	r := chi.NewRouter()
//
//	r.Route("/hello", func(r chi.Router) {
//		r.Get("/count", func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusOK)
//			json.NewEncoder(w).Encode(map[string]string{"Hello count": strconv.int(countOfRequest, 10)})
//		})
//	})
//
//	http.ListenAndServe(":8080", r)
//}
