package todo

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func NewTodo(name string) Todo {
	return Todo{
		Name:      name,
		Completed: false,
	}
}

type CreateTodo struct {
	Name string `json:"name"`
}
