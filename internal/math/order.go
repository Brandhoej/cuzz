package math

import "golang.org/x/exp/constraints"

func OrderOf[T constraints.Ordered]() func(lhs, rhs T) int {
	return func(lhs, rhs T) int {
		if lhs < rhs {
			return -1
		} else if lhs > rhs {
			return 1
		}
		return 0
	}
}
