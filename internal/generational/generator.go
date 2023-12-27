package generational

import (
	"context"
	"errors"
	"reflect"
)

var ErrGeneratorEmpty = errors.New("the generator cannot generate an element")

type Generator[T any] interface {
	Next(context context.Context) (T, error)
}

type RecursiveGenerator[T any] struct {
	blueprint any
	count int
}

func (generator RecursiveGenerator[T]) Next(context context.Context) (T, error) {
	panic("Not implemented")
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

func For[T any](blueprint any) Generator[T] {
	// If the specification already is a generator then we just return that generator.
	if generator, ok := blueprint.(Generator[T]); ok {
		return generator
	}

	// It is also possible to have specifications that are not generators.
	var target T
	concretion := reflect.ValueOf(target)
	specification := reflect.ValueOf(blueprint)
	if concretion.Kind() != specification.Kind() {
		panic("The specification and concretion kinds are not the same.")
	}

	switch specification.Kind() {
	case reflect.Bool, reflect.Uintptr, reflect.String,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		if value, ok := blueprint.(T); ok {
			return Sequence[T](value)
		}
	case reflect.Array:
		if values, ok := blueprint.([]T); ok {
			return Sequence[T](values...)
		}
	case reflect.Struct:
		/*cN := concretion.NumField()
		sN := specification.NumField()*/
	}

	panic("Not implemented yet")
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
