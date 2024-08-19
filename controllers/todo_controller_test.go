package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTodoService(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{
			name:       "number query",
			query:      "1",
			resultCode: http.StatusOK,
		},
		{
			name:       "string query",
			query:      "a",
			resultCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/v1/todo/%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/v1/todo/{todo_id:[0-9]+}", con.GetTodoHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: got %d but want %d\n", res.Code, tt.resultCode)
			}
		})
	}
}
