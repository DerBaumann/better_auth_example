package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"better_auth_example/internal/user"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Todos []Todo
}

func NewHandler() Handler {
	return Handler{Todos: []Todo{}}
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	_, err := user.FromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewEncoder(w).Encode(h.Todos); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	_, err := user.FromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var createTodo CreateTodo

	if err := json.NewDecoder(r.Body).Decode(&createTodo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(string(data))

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
