package generational

import (
	"github.com/brandhoej/cuzz/internal/arrangement"
	"github.com/brandhoej/cuzz/internal/math"
	"golang.org/x/exp/constraints"
)

// Boundary Value Analysis based on:
//   https://people.eecs.ku.edu/~saiedian/Teaching/814/Readings/Neates-BVA-Testing.pdf

type BoundaryStrategy[T any] func(
	first, last, avg T,
	proceed func(domain T) (T, bool),
	preceed func(domain T) (T, bool),
	order func(lhs, rhs T) int,
) []T

func BoundaryStrategyNumericInterval[T constraints.Integer | constraints.Float](
	interval Interval[T],
	step T,
	strategy BoundaryStrategy[T],
) []T {
	if interval.Empty() {
		return make([]T, 0)
	}

	first, last := interval.lower, interval.upper

	if interval.lowerOpen {
		first += step
	}

	if interval.upperOpen {
		last -= step
	}

	return BoundaryStrategyNumeric[T](
		first, last, step,
		math.MinOf[T](), math.MaxOf[T](),
		strategy,
	)
}

func BoundaryStrategyNumeric[T constraints.Integer | constraints.Float](
	first, last, step T,
	min, max T,
	strategy BoundaryStrategy[T],
) []T {
	return strategy(
		first, last, (first+last)/2,
		func(domain T) (T, bool) { return domain + step, domain <= max-step },
		func(domain T) (T, bool) { return domain - step, domain >= min+step },
		math.OrderOf[T](),
	)
}

// Boundary Value Analysis (BVA) calculates the boundary values for the range.
//
// The idea and motivation behind BVA is that errors tend to occur near the extremities of the input variables.
// The defects found on the boundaries of these input variables can obviously be the result of countless possibilities.
// But there are many common faults that result in errors more collated towards the boundaries of input variables.
//
// Example:
//
//	panic("Not implemented")
func BVA[D any](
	first, last, avg D,
	proceed func(domain D) (D, bool),
	preceed func(domain D) (D, bool),
	order func(lhs, rhs D) int,
) []D {
	equal := func(lhs, rhs D) bool {
		return order(lhs, rhs) == 0
	}
	less := func(lhs, rhs D) bool {
		return order(lhs, rhs) == -1
	}
	greater := func(lhs, rhs D) bool {
		return order(lhs, rhs) == 1
	}

	var values []D = make([]D, 0, 5)

	if equal(first, last) {
		values = append(values, first)
	} else {
		values = append(values, first, last)
	}

	if !(equal(first, avg) || equal(last, avg)) {
		values = append(values, avg)
	}

	if value, exists := proceed(first); exists && less(value, last) {
		values = append(values, value)
	}

	if value, exists := preceed(last); exists && greater(value, first) {
		values = append(values, value)
	}

	return values
}

// Robutsness Analysis (RA) calculates the BVA and the values immediately exceeding the extremes.
//
// The idea behind Robustness testing is to test for clean and dirty test cases.
//
//	Clean: Input variables that lie in the legitimate input range.
//	Dirty: Input variables that fall just outside this input domain.
//
// Example:
//
//	panic("Not implemented")
func RA[D any](
	first, last, avg D,
	proceed func(domain D) (D, bool),
	preceed func(domain D) (D, bool),
	order func(lhs, rhs D) int,
) []D {
	// BVA returns the "clean" values.
	values := BVA[D](first, last, avg, proceed, preceed, order)

	// Dirty preceeding value from "first".
	if value, exists := preceed(first); exists {
		values = append(values, value)
	}

	// Dirty proceeding value from "last".
	if value, exists := proceed(last); exists {
		values = append(values, value)
	}

	return values
}

// Multi Boundary Analysis (MBA) calculates the boundaries (Often BVA or RA) and takes the cartesian product to capture the dimensional extremes.
//
// Boundary Value analysis uses the critical fault assumption and therefore only tests for a single variable at a time assuming its extreme values.
// By disregarding this assumption we are able to test the outcome if more than one variable were to assume its extreme value.
// This is done by also including the cartesian product.
//
// Worst Case Analysis (WCA) calculates the BVA and takes the cartesian product to capture the dimensional extremes.
//
// BVA uses the critical fault assumption and therefore only tests for a single variable at a time assuming its extreme values.
// By disregarding this assumption we are able to test the outcome if more than one variable were to assume its extreme value.
// This is done by also including the cartesian product.
//
// Example (single dimension inputs):
//
//	panic("Not implemented")
//
// Robust Worst-Case Analysis (RWCA) calculates the RA and takes the cartesian product to capture the dimensional extremes.
//
// RA includes clean and dirty values and therefore tests both valid and invalid values.
// However, like BVA it assumes that the input dimensions are not correlated by only considering one variable at a time.
// This is fized, the same way as WCA, by including the cartesian product.
//
// Example (single dimension inputs):
//
//	panic("Not implemented")
func MBA[D any](strategies [](func() []D)) [][]D {
	var values [][]D = make([][]D, 0, len(strategies))
	for idx := range strategies {
		values = append(values, strategies[idx]())
	}
	return arrangement.Cartesian(values)
}
