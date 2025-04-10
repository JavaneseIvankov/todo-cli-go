package handlers_todo

import (
	"time"

	"github.com/javaneseivankov/todo-cli-go/utils/args_iterator"
	"github.com/javaneseivankov/todo-cli-go/utils/constants"
	"github.com/javaneseivankov/todo-cli-go/utils/time_utils"
)

type ShowPayload struct {
   Completed *bool
    DueBefore  *time.Time 
    DueAfter   *time.Time 
    Name   *string    
	 Overdue *bool
}

func (p *ShowPayload) bind(args *args_iterator.ArgsIterator) error {

  getTime := func (args *args_iterator.ArgsIterator) *time.Time  {
    t, _err := time_utils.ParseUserTime(args.GetNext())
    if _err != nil {
      err = _err
      return nil
    }
    return t
  }

	flags := map[string]func() {
    "-done": func() {v := true; p.Completed = &v},
    "-undone": func() {v := false; p.Completed = &v},
	 "-before": func() {p.DueBefore = getTime(args)},
	 "-after": func() {p.DueAfter = getTime(args)},
	 "-overdue": func() {v:= true; p.Overdue = &v},
	}

  for (args.HasNext()) {
    callback, exists := flags[args.GetNext()];
    if !exists {
      return constants.InvalidInputError
    }
    callback()
  }

  return nil
}


type  AddTodoPayload struct{
  name string
  due *time.Time
}

func (p *AddTodoPayload) bind(args *args_iterator.ArgsIterator) error {
	var err error;

  flags := map[string]func(){
    "-name": func() {p.name = args.GetNextContinous()},
    "-due": func() {p.due = args.GetNextTimeContinous(&err)},
  }

  for (args.HasNext()) {
    callback, exists := flags[args.GetNext()];
    if !exists {
      return constants.InvalidInputError
    }
    callback()
  }

  return nil
}

type ModifyTodoPayload struct {
  id int
  name string
  due *time.Time
}

func (p *ModifyTodoPayload) bind(args *args_iterator.ArgsIterator) error {
  var err error;

  flags := map[string]func(){
    "-id": func() {p.id = args.GetNextInt(&err)},
    "-name": func() {p.name = args.GetNextContinous()},
    "-due": func() {p.due = args.GetNextTimeContinous(&err)},
  }

  if err != nil {
    return constants.InvalidInputError
  }

  for (args.HasNext()) {
    callback, exists := flags[args.GetNext()];
    if !exists {return constants.InvalidInputError}
    callback()
  }

  return nil
}

