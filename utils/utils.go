package utils

import (
	"strconv"
)

func KeyExists[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

func ParseU32(input string, error *error) uint32 {
	id, err := strconv.ParseUint(input, 10, 32)
	error = &err
  return uint32(id)
}
