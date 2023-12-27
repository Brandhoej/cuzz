package generational

import (
	"context"
	"testing"

	"golang.org/x/exp/constraints"
)

func abs[T constraints.Integer | constraints.Float](number T) T {
	if number < 0 {
		return -number
	}
	return number
}

func TestIntegerOnceExample(t *testing.T) {
	generator := Sequence[int](1, 2, 3)

	for test := range Stream(generator, context.Background()) {
		actual := abs(test)

		if test < 0 && -actual != test {
			t.Error("Expected abs of", test, "to be", -actual)
		} else if test >= 0 && actual != test {
			t.Error("Expected abs of", test, "to be", actual)
		}
	}
}

func TestStructExample(t *testing.T) {
	generator := For[struct {
		value int
	}](
		// Constant value.
		struct {
			value int
		}{
			value: 0,
		},
		// Generated value.
		struct {
			value Generator[int]
		}{
			value: Sequence[int](1, 2, 3),
		},
	)

	for test := range Stream(generator, context.Background()) {
		actual := abs(test.value)

		if test.value < 0 && -actual != test.value {
			t.Error("Expected abs of", test.value, "to be", -actual)
		} else if test.value >= 0 && actual != test.value {
			t.Error("Expected abs of", test.value, "to be", actual)
		}
	}
}
