package generational

import "context"

type Array[T any] struct {
	index int
	inner []T
}

func (array *Array[T]) Next(context context.Context) (T, error) {
	if len(array.inner) >= array.index {
		var zeroT T
		return zeroT, ErrGeneratorEmpty
	}

	element := array.inner[array.index]
	array.index += 1
	return element, nil
}
