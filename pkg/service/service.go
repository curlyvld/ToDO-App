
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
//go:generate mockery --name=TodoList --output=pkg/service/mocks --case=underscore
type TodoList interface {
	Create(userId int, list ToDoApp.TodoList) (int, error)
	GetAll(userId int) ([]ToDoApp.TodoList, error)
	GetById(userId, listId int) (ToDoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoApp.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item ToDoApp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]ToDoApp.TodoItem, error)
	GetById(userId, itemId int) (ToDoApp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input ToDoApp.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
