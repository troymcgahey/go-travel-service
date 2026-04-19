package service

import "github.com/troymcgahey/go-travel-service/internal/models"

type TodoService struct {
	todos  []models.Todo
	nextID int
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos:  []models.Todo{},
		nextID: 1,
	}
}

func (s *TodoService) List() []models.Todo {
	return s.todos
}

func (s *TodoService) Create(text string) models.Todo {
	todo := models.Todo{
		ID:   s.nextID,
		Text: text,
		Done: false,
	}
	s.todos = append(s.todos, todo)
	s.nextID++
	return todo
}
