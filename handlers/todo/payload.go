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
    "-done": func() {done := true; p.Completed = &done},
    "-undone": func() {undone := false; p.Completed = &undone},
	 "-before": func() {p.DueBefore = getTime(args)},
	 "-after": func() {p.DueAfter = getTime(args)},
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
  due string
}

func (p *AddTodoPayload) bind(args args_iterator.ArgsIterator) error {

  flags := map[string]func(){
    "-task": func() {p.name = args.GetNext()},
    "-due": func() {p.due = args.GetNext()},
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

  getTime := func (args *args_iterator.ArgsIterator) *time.Time  {
    t, _err := time_utils.ParseUserTime(args.GetNext())
    if _err != nil {
      err = _err
      return nil
    }
    return t
  }

  flags := map[string]func(){
    "-id": func() {p.id = args.GetNextInt(&err)},
    "-name": func() {p.name = args.GetNext()},
    "-due": func() {p.due = getTime(args)},
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

