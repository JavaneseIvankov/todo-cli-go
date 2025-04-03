package handlers_todo

import (
	"fmt"

	"github.com/javaneseivankov/todo-cli-go/repository"
	"github.com/javaneseivankov/todo-cli-go/utils/args_iterator"
)

var repo = repository.NewTodoRepoImpl()

func displayTodos(todos []repository.Todo) {
  fmt.Println("----Todos----");
  for i, todo := range todos {
    fmt.Printf("%d --- %s -- %s",i, todo.Name, todo.Due.GoString())
  }
}

func AddTodoHandler(args *args_iterator.ArgsIterator) {
  payload := &AddTodoPayload{}

  err := payload.bind(*args); if err != nil {
    fmt.Println(err)
    return
  }
  repo.AddTodo(payload.name, nil)

  todos, _ := repo.GetTodos()
  displayTodos(todos)
}

func DoneHandler(args *args_iterator.ArgsIterator) {
  var err error
  id := args.GetNextInt(&err)
  if err != nil {
    panic("DoneHandler: Invalid user input") 
  }
  if err = repo.CompleteTodo(id); err != nil{
    fmt.Println(err)
  }
}


func ModifyHandler(args *args_iterator.ArgsIterator) {
  payload := &ModifyTodoPayload{}

  if err := payload.bind(args); err != nil {
    fmt.Println(err)
    return
  }

  repo.ModifyTodo(payload.id, &payload.name, payload.due)
}

