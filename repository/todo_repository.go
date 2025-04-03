package repository

import (
	"errors"
	"time"

	"github.com/javaneseivankov/todo-cli-go/utils"
)

type Todo struct {
  Id int
  Name string
  Due time.Time
  Completed bool
}


type ITodoRepository interface {
  AddTodo(name string, due *time.Time)  (int, error)
  DeleteTodo(id int) error
  CompleteTodo(id int) error
  GetTodos() ([]Todo, error)
  ModifyTodo(id int, name *string, due *time.Time) error
}

type TodoRepoImpl struct {
  idCount int
  store map[int]*Todo
}

func NewTodoRepoImpl() ITodoRepository {
  return &TodoRepoImpl{idCount: 0, store: make(map[int]*Todo)}
}

func (r *TodoRepoImpl) AddTodo(name string, due *time.Time) (int, error) {
  r.idCount++
  newId := r.idCount
  r.store[newId] = &Todo{Id: newId, Name: name, Due: time.Now().AddDate(0, 0, 1)}
  return newId, nil 
}

func (r *TodoRepoImpl) DeleteTodo(id int) error {
  delete(r.store, id)
  return nil
}

func (r *TodoRepoImpl) CompleteTodo(id int) error {
  if (utils.KeyExists(r.store, id)) {
    r.store[id].Completed = true
    return nil
  }
  return errors.New("Todo with corresponding id doesnt exist!!") 
}

func (r* TodoRepoImpl) GetTodos() ([]Todo, error) {
  todos := make([]Todo, 0, len(r.store))
  for _,todo := range r.store {
    todos = append(todos, *todo)
  }
  return todos, nil
}

func (r *TodoRepoImpl) ModifyTodo(id int, name *string, due *time.Time) error {
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

