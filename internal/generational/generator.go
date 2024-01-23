package generational

import (
	"context"
	"errors"
)

var (
	ErrGeneratorEmpty = errors.New("the generator cannot generate an element")
)

type Generator[T any] interface {
	Next(context context.Context) (T, error)
}

type SequenceGenerator[T any] struct {
	index int
	inner []T
}

func (array *SequenceGenerator[T]) Next(context context.Context) (T, error) {
	if array.index >= len(array.inner) {
		var zeroT T
		return zeroT, ErrGeneratorEmpty
	}

	element := array.inner[array.index]
	array.index += 1
	return element, nil
}

func Boolean() Generator[bool] {
	return Sequence[bool](true, false)
}

func Sequence[T any](values ...T) Generator[T] {
	return &SequenceGenerator[T]{
		index: 0,
		inner: values,
	}
}

func Stream[T any](generator Generator[T], context context.Context) <-chan T {
	iterator := make(chan T)

	go func() {
		defer close(iterator)

		for {
			value, err := generator.Next(context)
			if err != nil {
				return
			}

			select {
			case <-context.Done():
				return
			case iterator <- value:
			}
		}
	}()

	return iterator
}