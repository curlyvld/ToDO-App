package repository

import (
	"ToDoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ToDoApp.User) (int, error)
	GetUser(username, password string) (ToDoApp.User, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
