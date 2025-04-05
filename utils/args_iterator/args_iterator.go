package args_iterator

import (
	"strconv"
	"time"

	"github.com/javaneseivankov/todo-cli-go/utils/time_utils"
)

type ArgsIterator struct {
  args []string
  Index int 
  valid bool
}

func isFlag(input string) bool {
	if (input[0] == '-') {
		return true
	}
	return false
}

func NewArgsIterator (args []string) *ArgsIterator {
  return &ArgsIterator{
    args: args,
    valid: true,
  }
}


func (i *ArgsIterator) HasNext() bool {
  if (i.Index < len(i.args)) {
    return true
  }
  return false
}

func (i *ArgsIterator) GetNext() string {
  if i.HasNext() {
    arg:=i.args[i.Index]
    i.Index++
    return arg
  }
  return ""
}

func (i *ArgsIterator) GetNextContinous() string {
	ret := ""
	first := false
	for i.HasNext() {
		if isFlag(i.args[i.Index]) {
		 return ret
		}
		if (!first) {ret += " "}
		ret += i.args[i.Index]
		i.Index++
	}
	return ret
}

func (i *ArgsIterator) GetNextInt(error *error) int {
  next, err := strconv.Atoi(i.GetNext())
  error = &err
  return next
}

func  (i *ArgsIterator) GetNextTime(error *error) *time.Time  {
	t, _err := time_utils.ParseUserTime(i.GetNext())
	if _err != nil {
	error = &_err
	return nil
	}
	return t
  }

func (i* ArgsIterator) GetNextTimeContinous(error *error) *time.Time {
	t, _err := time_utils.ParseUserTime(i.GetNextContinous())
	if _err != nil {
	error = &_err
	return nil
	}
	return t
}
