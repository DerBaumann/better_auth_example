package todo

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Todos []Todo
}

func NewHandler() Handler {
	return Handler{Todos: []Todo{}}
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(h.Todos); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodo CreateTodo

	if err := json.NewDecoder(r.Body).Decode(&createTodo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	todo := NewTodo(createTodo.Name)
	h.Todos = append(h.Todos, todo)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetTodos)
	r.Post("/", h.CreateTodo)

	return r
}
