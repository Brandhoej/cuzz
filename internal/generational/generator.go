package generational

import (
	"context"
	"errors"
)

var ErrGeneratorEmpty = errors.New("the generator cannot generate an element")

type Generator[T any] interface {
	Next(context context.Context) (T, error)
}

func Boolean() Generator[bool] {
	return Once[bool](true, false)
}

func Once[T any](values ...T) Generator[T] {
	return &Array[T]{
		index: 0,
		inner: values,
	}
}
