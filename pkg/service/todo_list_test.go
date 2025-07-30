package service

import (
	"testing"

	"ToDoApp"
	"ToDoApp/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
)

// 2. Пишем сам тест
func TestTodoListService_Create(t *testing.T) {
	// Подготавливаем входные данные
	userId := 1
	list := ToDoApp.TodoList{
		Title:       "Test List",
		Description: "Test Description",
	}
	returnedId := 42 // то, что должен вернуть мок

	// Создаём мок-репозиторий и настраиваем ожидания
	mockRepo := new(mocks.TodoList)
	mockRepo.On("Create", userId, list).Return(returnedId, nil)

	// Создаём сервис с этим мок-репозиторием
	service := NewTodoListService(mockRepo)

	// Вызываем тестируемый метод
	id, err := service.Create(userId, list)

	// Проверяем результат
	assert.NoError(t, err) // Ожидаем, что ошибки не будет
	assert.Equal(t, returnedId, id) // Ожидаем, что id совпадёт

	// Проверяем, что мок был вызван с нужными аргументами
	mockRepo.AssertExpectations(t)
}

// Пояснения к шагам:
// 1. MockTodoListRepo реализует только нужный метод Create.
// 2. В тесте мы задаём, что при вызове Create с определёнными аргументами мок должен вернуть (returnedId, nil).
// 3. Создаём сервис с этим мок-репозиторием.
// 4. Вызываем Create и проверяем, что результат совпадает с ожидаемым.
// 5. AssertExpectations проверяет, что мок был вызван так, как мы ожидали. 