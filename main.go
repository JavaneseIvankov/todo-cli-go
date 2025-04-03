package main

import (
	"fmt"
	"os"

	handlers_todo "github.com/javaneseivankov/todo-cli-go/handlers/todo"
	"github.com/javaneseivankov/todo-cli-go/utils/args_iterator"
)


func main() {
  args := args_iterator.NewArgsIterator(os.Args[1:]);
  if args.HasNext() {
    switch(args.GetNext()) {
    case "add":
      handlers_todo.AddTodoHandler(args);
      break;
    case "done":
      handlers_todo.DoneHandler(args);
      break
    case "modify":
      handlers_todo.ModifyHandler(args)
      break;
    default:
      fmt.Println("Invalid input!")
      break;
  }
 }
}
