package main

import (
	"github.com/javaneseivankov/todo-cli-go/utils/args_iterator"
	"github.com/javaneseivankov/todo-cli-go/utils/constants"
)

type HandleFunc func(...string) error

type Handler struct {
  args args_iterator.ArgsIterator
  paramsCount int
  handlerFunc HandleFunc
}

func (h *Handler) getParams() ([]string, error) {
  var params []string;
  for range h.paramsCount {
    if (!h.args.HasNext() && len(params) < h.paramsCount) {
      return nil, constants.InvalidInputError
    }
    params = append(params, h.args.GetNext())
  }
  return params, nil
}

func (h *Handler) Handle() error {
  params, err := h.getParams();
  if err != nil {
    panic("lol")
  }
  if err := h.handlerFunc(params...); err != nil {
    panic ("lol2")
  }
  return nil
}
