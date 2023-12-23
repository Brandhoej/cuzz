package ebnf

type Queue[T any] interface {
	Size() int
	Enqueue(element T)
	Dequeue() T
	Peek() T
}

type ArrayQueue[T any] struct {
	array []T
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		array: make([]T, 0),
	}
}

func (queue *ArrayQueue[T]) Size() int {
	return len(queue.array)
}

func (queue *ArrayQueue[T]) Enqueue(element T) {
	queue.array = append(queue.array, element)
}

func (queue *ArrayQueue[T]) Dequeue() (element T) {
	element = queue.array[0]
	queue.array = queue.array[1:]
	return
}

func (queue *ArrayQueue[T]) Peek() T {
	return queue.array[0]
}
