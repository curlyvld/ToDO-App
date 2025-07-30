package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ToDoApp"
	"ToDoApp/pkg/service"
	"ToDoApp/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"errors"
)

// Вспомогательная функция для подмены userId в контексте
func addUserIdToContext(c *gin.Context, userId int) {
	c.Set("userId", userId)
}

func TestHandler_createList(t *testing.T) {
	// 1. Подготавливаем входные данные
	userId := 123
	input := ToDoApp.TodoList{
		Title:       "Test List",
		Description: "Test Description",
	}
	returnedId := 42

	// 2. Создаём мок-сервис
	mockTodoList := new(mocks.TodoList)
	mockTodoList.On("Create", userId, input).Return(returnedId, nil)

	// Используем реальную структуру service.Service
	services := &service.Service{
		TodoList: mockTodoList,
	}
	h := &Handler{services: services}

	// 4. Создаём gin-контекст и http-запрос
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/lists", func(c *gin.Context) {
		addUserIdToContext(c, userId) // подставляем userId
		h.createList(c)
	})

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/lists", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// 5. Выполняем запрос
	r.ServeHTTP(w, req)

	// 6. Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, float64(returnedId), resp["id"]) // json.Unmarshal всегда даёт float64 для чисел

	// 7. Проверяем, что мок был вызван
	mockTodoList.AssertExpectations(t)
}

// Тест для получения всех списков дел (GET /lists)
func TestHandler_getAllLists(t *testing.T) {
	// 1. Подготавливаем входные данные
	userId := 123
	mockLists := []ToDoApp.TodoList{
		{
			Id:          1,
			Title:       "Work Tasks",
			Description: "Tasks for work",
		},
		{
			Id:          2,
			Title:       "Personal Tasks",
			Description: "Personal todo items",
		},
	}

	// 2. Создаём мок-сервис и настраиваем его поведение
	mockTodoList := new(mocks.TodoList)
	// Ожидаем, что при вызове GetAll с userId=123 вернётся наш список
	mockTodoList.On("GetAll", userId).Return(mockLists, nil)

	// 3. Создаём сервисы с моком
	services := &service.Service{
		TodoList: mockTodoList,
	}
	h := &Handler{services: services}

	// 4. Настраиваем gin и маршрут
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/lists", func(c *gin.Context) {
		// Подставляем userId в контекст (эмулируем авторизацию)
		c.Set("userId", userId)
		h.getAllLists(c)
	})

	// 5. Создаём HTTP-запрос
	req, _ := http.NewRequest(http.MethodGet, "/lists", nil)
	w := httptest.NewRecorder()

	// 6. Выполняем запрос
	r.ServeHTTP(w, req)

	// 7. Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code) // Ожидаем статус 200

	// Парсим JSON-ответ
	var response getAllListsResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Проверяем, что вернулись правильные данные
	assert.Len(t, response.Data, 2) // Ожидаем 2 списка
	assert.Equal(t, mockLists[0].Title, response.Data[0].Title)
	assert.Equal(t, mockLists[1].Title, response.Data[1].Title)

	// 8. Проверяем, что мок был вызван с правильными аргументами
	mockTodoList.AssertExpectations(t)
}

// Тест для обработки ошибки при получении списков
func TestHandler_getAllLists_Error(t *testing.T) {
	userId := 123
	expectedError := "database error"

	// Настраиваем мок, чтобы он возвращал ошибку
	mockTodoList := new(mocks.TodoList)
	mockTodoList.On("GetAll", userId).Return([]ToDoApp.TodoList{}, errors.New(expectedError))

	services := &service.Service{
		TodoList: mockTodoList,
	}
	h := &Handler{services: services}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/lists", func(c *gin.Context) {
		c.Set("userId", userId)
		h.getAllLists(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/lists", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Проверяем, что вернулся статус ошибки
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockTodoList.AssertExpectations(t)
} 