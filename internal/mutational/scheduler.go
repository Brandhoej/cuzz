package mutational

import "math/rand"

type Schedule[T any] []Operator[T]

type Scheduler[T any] interface {
	Schedule(seed T) (Schedule[T], error)
}

type UniformScheduler[T any] struct {
	operators []Operator[T]
	amount    int
}

func (scheduler UniformScheduler[T]) Schedule(
	seed T,
) (Schedule[T], error) {
	length := len(scheduler.operators)
	schedule := make(Schedule[T], scheduler.amount)
	for i := 0; i < scheduler.amount; i++ {
		schedule[i] = scheduler.operators[rand.Intn(length)]
	}
	return schedule, nil
}
