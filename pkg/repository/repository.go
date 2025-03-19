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
	Create(listId int, item ToDoApp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]ToDoApp.TodoItem, error)
	GetById(userId, itemId int) (ToDoApp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input ToDoApp.UpdateItemInput) error
}

type TodoList interface {
	Create(userId int, list ToDoApp.TodoList) (int, error)
	GetAll(userId int) ([]ToDoApp.TodoList, error)
	GetById(userId, listId int) (ToDoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoApp.UpdateListInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
