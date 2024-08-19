package services

import (
	"database/sql"
	"errors"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/models"
	"github.com/AppDeveloperMLLB/todo_app/repositories"
)

// GetTodoService - get todo
func (s *MyAppService) GetTodoService(todoID int) (models.Todo, error) {
	todo, err := repositories.SelectTodo(s.db, todoID)
	if errors.Is(err, sql.ErrNoRows) {
		err = apperrors.NAData.Wrap(err, "no data")
		return models.Todo{}, err
	}

	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Todo{}, err
	}

	return todo, nil
}

// GetTodoListService - get todo list
func (s *MyAppService) GetTodoListService(page int, perPage int) ([]models.Todo, error) {
	list, err := repositories.SelectTodoList(s.db, page, perPage, "all")
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(list) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return list, nil
}

// CreateTodoService - create todo
func (s *MyAppService) CreateTodoService(todo models.Todo) (models.Todo, error) {
	createdTodo, err := repositories.CreateTodo(s.db, todo)
	if err != nil {
		err = apperrors.InsertDataField.Wrap(err, "fail to insert data")
		return models.Todo{}, err
	}

	return createdTodo, nil
}

// UpdateTodoService
func (s *MyAppService) UpdateTodoService(todo models.Todo) (models.Todo, error) {
	updatedTodo, err := repositories.UpdateTodo(s.db, todo)
	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return models.Todo{}, err
	}

	return updatedTodo, nil
}

// DeleteTodoService
func (s *MyAppService) DeleteTodoService(todoID int, userID string) error {
	todo, err := s.GetTodoService(todoID)
	if err != nil {
		return err
	}

	if todo.UserID != userID {
		err = apperrors.Forbidden.Wrap(err, "forbidden")
		return err
	}

	err = repositories.DeleteTodo(s.db, todoID)
	if err != nil {
		err = apperrors.DeleteDataFailed.Wrap(err, "fail to delete data")
		return err
	}

	return nil
}
