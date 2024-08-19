package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/AppDeveloperMLLB/todo_app/models"
	"github.com/AppDeveloperMLLB/todo_app/repositories"
)

func TestSelectTodo(t *testing.T) {
	type expected struct {
		todo models.Todo
		err  error
	}
	tests := []struct {
		testTitle string
		todoID    int
		expected  expected
	}{
		{
			testTitle: "subtest1",
			todoID:    1,
			expected: expected{
				todo: models.Todo{
					ID:          1,
					UserID:      "google-123",
					Title:       "Todo1dd",
					Description: "Fiction",
					Status:      "todo",
				},
				err: nil,
			},
		},
		{
			testTitle: "subtest2",
			todoID:    999,
			expected: expected{
				todo: models.Todo{},
				err:  sql.ErrNoRows,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectTodo(testDB, test.todoID)
			if err != test.expected.err {
				t.Errorf("error: get %v but want %v\n", err, test.expected.err)
			}

			if got.ID != test.expected.todo.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.todo.ID)
			}

			if got.UserID != test.expected.todo.UserID {
				t.Errorf("UserID: get %s but want %s\n", got.UserID, test.expected.todo.UserID)
			}

			if got.Title != test.expected.todo.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.todo.Title)
			}

			if got.Description != test.expected.todo.Description {
				t.Errorf("Description: get %s but want %s\n", got.Description, test.expected.todo.Description)
			}

			if got.Status != test.expected.todo.Status {
				t.Errorf("Status: get %s but want %s\n", got.Status, test.expected.todo.Status)
			}
		})
	}
}

// parameter test
func TestSelectTodoList(t *testing.T) {
	tests := []struct {
		testTitle string
		page      int
		perPage   int
		expected  []models.Todo
	}{
		{
			testTitle: "subtest1",
			page:      1,
			perPage:   2,
			expected: []models.Todo{
				{
					ID:          1,
					UserID:      "google-123",
					Title:       "Todo1",
					Description: "Fiction",
					Status:      "todo",
				},
				{
					ID:          2,
					UserID:      "google-123",
					Title:       "Todo2",
					Description: "Non-Fiction",
					Status:      "in_progress",
				},
			},
		},
		{
			testTitle: "subtest2",
			page:      2,
			perPage:   2,
			expected: []models.Todo{
				{
					ID:          3,
					UserID:      "google-123",
					Title:       "Todo3",
					Description: "Science Fiction",
					Status:      "completed",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			gotList, err := repositories.SelectTodoList(testDB, test.page, test.perPage, "all")
			if err != nil {
				t.Fatal(err)
			}

			if len(gotList) != len(test.expected) {
				t.Errorf("len: get %d but want %d\n", len(gotList), len(test.expected))
			}

			for i, v := range gotList {
				got := v
				expected := test.expected[i]

				if got.ID != expected.ID {
					t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
				}
				if got.UserID != expected.UserID {
					t.Errorf("UserID: get %s but want %s\n", got.UserID, expected.UserID)
				}
				if got.Title != expected.Title {
					t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
				}
				if got.Description != expected.Description {
					t.Errorf("Description: get %s but want %s\n", got.Description, expected.Description)
				}
				if got.Status != expected.Status {
					t.Errorf("Status: get %s but want %s\n", got.Status, expected.Status)
				}
			}
		})
	}
}
