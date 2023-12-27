package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

func SafeUnsignedDifference[S constraints.Signed, U constraints.Unsigned](lhs, rhs S) U {
	var min, max S = lhs, rhs
	if min > max {
		min, max = max, min
	}

	if min <= 0 && max <= 0 {
		// E.g., min=-10, max=-6 -> 10 - 6
		return U(-min) - U(-max)
	} else if min >= 0 && max >= 0 {
		// E.g., min=1, max=4 -> 4 - 1
		return U(max) - U(min)
	}

	// E.g., min=-2, max=3 -> 2 + 3
	return U(-min) + U(max)
}

func SafeUnsignedAddition[S constraints.Signed, U constraints.Unsigned](lhs S, rhs U) (val S) {
	max := U(MaxOf[S]())

	val = lhs
	rest := U(rhs)

	for rest > 0 {
		if rest > max {
			val += S(max)
			rest -= max
		} else if rest > 0 {
			val += S(rest)
			rest = 0
		}
	}

	return
}

func Lerp[T constraints.Float](from, to, t T) T {
	return from*(1.0-t) + (to * t)
}

func IsInteger[T constraints.Integer | constraints.Float]() bool {
	var value T
	switch any(value).(type) {
	case float32, float64:
		return false
	}
	return true
}

func SmallestNonZero[T constraints.Integer | constraints.Float]() (value T) {
	switch any(value).(type) {
	case float32:
		var v float32 = math.SmallestNonzeroFloat32
		value = T(v)
	case float64:
		var v float32 = math.SmallestNonzeroFloat64
		value = T(v)
	default:
		value = 1
	}
	return
}
