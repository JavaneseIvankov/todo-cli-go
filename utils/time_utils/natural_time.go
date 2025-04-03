package time_utils

import (
	"errors"
	"time"

	"github.com/tj/go-naturaldate"
)

func ParseUserTime(input string) (*time.Time, error) {
  ref := time.Now()
  t, err := naturaldate.Parse(input, ref) 
  if err != nil {
    return nil, errors.New("Failed parsing time from user input")
  }
  return &t, nil
}
