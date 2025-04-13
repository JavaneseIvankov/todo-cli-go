package handlers_todo

import (
	"fmt"
	"strconv"

	repository "github.com/javaneseivankov/todo-cli-go/repository/todo_repository"
	"github.com/javaneseivankov/todo-cli-go/ui/cli"
	"github.com/javaneseivankov/todo-cli-go/utils/args_iterator"
	"github.com/javaneseivankov/todo-cli-go/utils/time_utils"
)

var repo, err = repository.NewSQLiteTodoRepo("todo.db")

// func displayTodos(todos []repository.Todo) {
//   fmt.Println()
//   fmt.Println("----Todos----");
//   fmt.Printf("%s\t %s\t %s\t %s\n", "Id", "Name", "Due", "Completed");
//   for _, todo := range todos {
//     fmt.Printf("%d\t %s\t %s\t %s\n", todo.Id, todo.Name, time_utils.ToHumandReadable(todo.Due),  time_utils.ToHumandReadable(*todo.Completed))
//   }
//   fmt.Println()
// }

func getDisplayValues(todo repository.Todo) []string {
	id, name, due, completed := "-", "-", "-", "-"
	id = strconv.Itoa(todo.Id)
	name = todo.Name
	due = time_utils.ToHumandReadable(todo.Due)
	if todo.Completed != nil {
		completed = time_utils.ToHumandReadable(*todo.Completed)
	}
	return []string{id, name, due, completed}
}

func displayTodos(todos []repository.Todo) {
	table := cli.CreateTable(4)
  table.SetHeaders([]string{"ID", "Name", "Due", "Completed"})
	for _, todo := range todos {
		values := getDisplayValues(todo)
		table.Append(values)
	}
	table.Render()
}

func AddTodoHandler(args *args_iterator.ArgsIterator) {
  payload := &AddTodoPayload{}

  err := payload.bind(args); if err != nil {
    fmt.Println(err)
    return
  }
  repo.AddTodo(payload.name, payload.due)

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

func ShowHandler(args *args_iterator.ArgsIterator) {
   payload := &ShowPayload{}

  if err := payload.bind(args); err != nil {
    fmt.Println(err)
    return
  }

  // This is bad, should've been in service layer
  // i need to separate the payload and query filter to fix circular import quickly
  filter := repository.QueryFilter{
   Completed: payload.Completed,
   DueBefore: payload.DueBefore,
   DueAfter: payload.DueAfter,
   NameLike: payload.Name,
   Overdue: payload.Overdue,
  }

  todos, err := repo.GetTodos(filter);
  if err != nil {
   fmt.Println(err)
  }

  displayTodos(todos)
}