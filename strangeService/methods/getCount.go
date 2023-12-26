package methods

import (
	"fmt"
	"github.com/google/uuid"
)

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

var Tasks []Task

type Task struct {
	ID        uuid.UUID
	Desc      string
	Completed bool
}

func newTask(desc string) int {
	task := Task{
		ID:        uuid.New(),
		Desc:      desc,
		Completed: false,
	}

	Tasks = append(Tasks, task)

	return len(Tasks) - 1
}

func main() {
	id := newTask("ПП работает фигово")
	fmt.Println(Tasks[id])

	changeStatus(id)
	fmt.Println(Tasks[id])
}

func changeStatus(id int) {
	Tasks[id].Completed = true
}
