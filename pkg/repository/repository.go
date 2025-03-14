package repository

import "github.com/jmoiron/sqlx"

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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
