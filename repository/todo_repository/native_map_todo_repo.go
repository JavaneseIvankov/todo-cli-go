package repository

import (
	"errors"
	"time"

	"github.com/javaneseivankov/todo-cli-go/utils"
)

type NativeMapTodoRepo struct {
  idCount int
  store map[int]*Todo
}

func NewNativeMapTodoRepo() ITodoRepository {
  return &NativeMapTodoRepo{idCount: 0, store: make(map[int]*Todo)}
}

func (r *NativeMapTodoRepo) AddTodo(name string, due *time.Time) (int, error) {
  r.idCount++
  newId := r.idCount
  r.store[newId] = &Todo{Id: newId, Name: name, Due: time.Now().AddDate(0, 0, 1)}
  return newId, nil 
}

func (r *NativeMapTodoRepo) DeleteTodo(id int) error {
  delete(r.store, id)
  return nil
}

func (r *NativeMapTodoRepo) CompleteTodo(id int) error {
  if (utils.KeyExists(r.store, id)) {
    r.store[id].Completed = true
    return nil
  }
  return errors.New("Todo with corresponding id doesnt exist!!") 
}

func (r* NativeMapTodoRepo) GetTodos(filter QueryFilter) ([]Todo, error) {
  todos := make([]Todo, 0, len(r.store))
  for _,todo := range r.store {
    todos = append(todos, *todo)
  }
  return todos, nil
}

func (r *NativeMapTodoRepo) ModifyTodo(id int, name *string, due *time.Time) error {
  if (utils.KeyExists(r.store, id)) {
      todo := r.store[id]

      if name != nil {
        todo.Name = *name
      }
      if due != nil {
        todo.Due = *due
      
      r.store[id] = todo 
      return nil

    }
}
  return errors.New("Todo with corresponding id not found")
}
