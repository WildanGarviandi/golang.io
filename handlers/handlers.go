package handlers

import(
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
  "api.etobee.com/api/model"
)

func Index(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, "Welcome! \n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  todos := model.Todos{
		model.Todo{Name: "Write presentation"},
		model.Todo{Name: "Host meetup"},
    model.Todo{Name: "Haha this is example 3", Completed: true},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprint(w, "Todo show : %s \n", todoId)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  todos := model.Todos{
		model.Todo{Name: "Write presentation"},
		model.Todo{Name: "Host meetup"},
    model.Todo{Name: "Haha this is example 3", Completed: true},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
