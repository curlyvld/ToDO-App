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
	Create(userId int, list ToDoApp.TodoList) (int, error)
	GetAll(userId int) ([]ToDoApp.TodoList, error)
	GetById(userId, listId int) (ToDoApp.TodoList, error)
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
