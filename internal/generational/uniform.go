package generational

import (
	"context"
	"math/rand"

	"github.com/brandhoej/cuzz/internal/arbitrary"
	"golang.org/x/exp/constraints"
)

type Uniform[T any] interface {
	Next() T
}

type UniformInterval[T constraints.Integer | constraints.Float] struct {
	interval Interval[T]
	step     T
	prng     *rand.Rand
}

func (uniform UniformInterval[Integral]) lower() Integral {
	var lower Integral
	if lower, open := uniform.interval.Lower(); open {
		lower += uniform.step
	}
	return lower
}

func (uniform UniformInterval[Integral]) upper() Integral {
	var upper Integral
	if upper, open := uniform.interval.Upper(); open {
		upper -= uniform.step
	}
	return upper
}

func (uniform UniformInterval[T]) Next() T {
	return arbitrary.InRange[T](uniform.prng, uniform.lower(), uniform.upper())
}

type UniformGenerator[T any] struct {
	uniform Uniform[T]
}

func (generator *UniformGenerator[T]) Next(_ context.Context) (T, error) {
	return generator.uniform.Next(), nil
}
