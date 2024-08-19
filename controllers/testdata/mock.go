package testdata

import "github.com/AppDeveloperMLLB/todo_app/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

type TodoService interface {
	GetTodoService(todoID int) (models.Todo, error)
	GetTodoListService(page int, perPage int) ([]models.Todo, error)
	CreateTodoService(todo models.Todo) (models.Todo, error)
	UpdateTodoService(todo models.Todo) (models.Todo, error)
	DeleteTodoService(todoID int, userID string) error
}

func (s *serviceMock) GetTodoService(todoID int) (models.Todo, error) {
	return todoDataList[0], nil
}

func (s *serviceMock) GetTodoListService(page int, perPage int) ([]models.Todo, error) {
	return todoDataList, nil
}

func (s *serviceMock) CreateTodoService(todo models.Todo) (models.Todo, error) {
	return todoDataList[0], nil
}

func (s *serviceMock) UpdateTodoService(todo models.Todo) (models.Todo, error) {
	return todoDataList[0], nil
}

func (s *serviceMock) DeleteTodoService(todoID int, userID string) error {
	return nil
}
