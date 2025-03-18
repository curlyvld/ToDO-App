package service

import (
	"ToDoApp"
	"ToDoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user ToDoApp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list ToDoApp.TodoList) (int, error)
	GetAll(userId int) ([]ToDoApp.TodoList, error)
	GetById(userId, listId int) (ToDoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoApp.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
