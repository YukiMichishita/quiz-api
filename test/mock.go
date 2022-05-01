package test

import (
	"errors"
	"net/http"
)

type MockTodoController struct {
}

func (mtc *MockTodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}

func (mtc *MockTodoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func (mtc *MockTodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

type MockTodoRepository struct {
}

func (mtr *MockTodoRepository) GetTodos() (todos []model.Todomodel, err error) {
	todos = []model.Todomodel{}
	return
}

func (mtr *MockTodoRepository) InsertTodo(todo model.Todomodel) (id int, err error) {
	id = 2
	return
}

func (mtr *MockTodoRepository) UpdateTodo(todo model.Todomodel) (err error) {
	return
}

func (mtr *MockTodoRepository) DeleteTodo(id int) (err error) {
	return
}

type MockTodoRepositoryGetTodosExist struct {
}

func (mtrgex *MockTodoRepositoryGetTodosExist) GetTodos() (todos []model.Todomodel, err error) {
	todos = []model.Todomodel{}
	todos = append(todos, model.Todomodel{Id: 1, Title: "title1", Content: "contents1"})
	todos = append(todos, model.Todomodel{Id: 2, Title: "title2", Content: "contents2"})
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) InsertTodo(todo model.Todomodel) (id int, err error) {
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) UpdateTodo(todo model.Todomodel) (err error) {
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) DeleteTodo(id int) (err error) {
	return
}

type MockTodoRepositoryError struct {
}

func (mtrgtn *MockTodoRepositoryError) GetTodos() (todos []model.Todomodel, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgie *MockTodoRepositoryError) InsertTodo(todo model.Todomodel) (id int, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgue *MockTodoRepositoryError) UpdateTodo(todo model.Todomodel) (err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgde *MockTodoRepositoryError) DeleteTodo(id int) (err error) {
	err = errors.New("unexpected error occurred")
	return
}
