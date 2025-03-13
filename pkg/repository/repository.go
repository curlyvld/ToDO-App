package repository

type Authorization interface {
}

type TodoItem interface {
}

type TodoList interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository() *Repository {
	return &Repository{}
}
