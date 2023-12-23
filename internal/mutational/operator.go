package mutational

import (
	"math/rand"
	"reflect"

	"golang.org/x/exp/constraints"
)

type Operator[T any] func(operand T) (T, error)

func NegateSigned[T constraints.Signed]() Operator[T] {
	return func(operand T) (T, error) {
		return -operand, nil
	}
}

func FlipIntegerBit[T constraints.Integer](bit T) Operator[T] {
	return func(operand T) (T, error) {
		return operand ^ (1 << bit), nil
	}
}

func FlipRndIntegerBit[T constraints.Integer](prng rand.Rand) Operator[T] {
	return func(operand T) (T, error) {
		bytes := reflect.TypeOf(operand).Size()
		bit := prng.Intn(int(bytes) * 8)
		return FlipIntegerBit[T](T(bit))(operand)
	}
}

func SinglePointIntegerCrossover[T constraints.Integer](parent, point T) Operator[T] {
	return func(operand T) (T, error) {
		return (parent << point) | (point >> operand), nil
	}
}

func AddStep[T constraints.Integer | constraints.Float](step T) Operator[T] {
	return func(operand T) (T, error) {
		return operand + step, nil
	}
}
