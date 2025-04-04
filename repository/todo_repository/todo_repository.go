package repository

import (
	"time"
)

type Todo struct {
  Id int
  Name string
  Due time.Time
  Completed bool
}

type QueryFilter struct {
    Completed  *bool
    DueBefore  *time.Time 
    DueAfter   *time.Time 
    NameLike   *string    
    Overdue    *bool
}



type ITodoRepository interface {
  AddTodo(name string, due *time.Time)  (int, error)
  DeleteTodo(id int) error
  CompleteTodo(id int) error
  GetTodos(filter QueryFilter) ([]Todo, error)
  ModifyTodo(id int, name *string, due *time.Time) error
}
