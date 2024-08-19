package controllers_test

import (
	"testing"

	"github.com/AppDeveloperMLLB/todo_app/controllers"
	"github.com/AppDeveloperMLLB/todo_app/controllers/testdata"
)

// sut
var con *controllers.TodoController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	con = controllers.NewTodoController(ser)

	m.Run()
}
