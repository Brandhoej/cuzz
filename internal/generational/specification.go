package generational

type Blueprint[T any] struct {
	specification any
}

func (blueprint Blueprint[T]) Generator() Generator[T] {
	panic("Not implemented yet")
}
