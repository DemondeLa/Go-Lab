package slice

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("index out of range")

func Delete[T any](s []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(s) {
		return s,
			fmt.Errorf("%w: index %d, length %d", ErrIndexOutOfRange, idx, len(s))
	}

	copy(s[idx:], s[idx+1:])

	var zero T
	s[len(s)-1] = zero

	return s[:len(s)-1], nil
}

func Add[T any](s []T, idx int, val T) ([]T, error) {
	if idx < 0 || idx > len(s) {
		return s,
			fmt.Errorf("%w: index %d, length %d", ErrIndexOutOfRange, idx, len(s))
	}

	var zero T
	s = append(s, zero)
	copy(s[idx+1:], s[idx:])
	s[idx] = val

	return s, nil
}
