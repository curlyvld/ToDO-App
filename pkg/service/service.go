package service

import "ToDoApp/pkg/repository"

type Authorization interface {
}

type TodoItem interface {
}

type TodoList interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
