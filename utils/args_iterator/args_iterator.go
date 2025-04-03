package args_iterator

import "strconv"

type ArgsIterator struct {
  args []string
  Index int 
  valid bool
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

func (i *ArgsIterator) GetNextInt(error *error) int {
  next, err := strconv.Atoi(i.GetNext())
  error = &err
  return next
}
