package generational

import (
	"cmp"
	"errors"
	"sort"
)

var (
	ErrNoIntersection = errors.New("intervals do not intersect")
	ErrNotExtendable  = errors.New("intervals are not extendable")
	ErrNotSorted      = errors.New("intervals are not sorted")
)

/* Represents an interval with comparable extreme values.
 * Closed Interval: If an interval is closed at a particular endpoint,
 *   it means that the value at that endpoint is included in the interval.
 *   Mathematically, for a closed interval [a, b], both a and b are part of the interval.
 * Open Interval: If an interval is open at a particular endpoint,
 *   it means that the value at that endpoint is not included in the interval.
 *   Mathematically, for an open interval (a, b), neither a nor b are part of the interval.
 *
 * Closed Interval: Includes both endpoints. Example: [1, 5] includes 1 and 5.
 * Open Interval: Excludes both endpoints. Example: (1, 5) excludes 1 and 5.
 * Half-Open or Half-Closed Interval: One endpoint is included, and the other is excluded.
 *   Example: [1, 5) includes 1 but excludes 5.
 *   Example: (1, 5] excludes 1 but includes 5.
 *
 * An interval is empty if it is closed and the extremes are the same.
 */
type Interval[T cmp.Ordered] struct {
	lowerOpen bool
	upperOpen bool
	lower     T
	upper     T
}

func Create[T cmp.Ordered](
	lower, upper T,
	lowerOpen, upperOpen bool,
) Interval[T] {
	return Interval[T]{
		lowerOpen,
		upperOpen,
		lower,
		upper,
	}
}

func (interval Interval[T]) Empty() bool {
	return interval.lowerOpen && interval.upperOpen &&
		interval.lower == interval.upper
}

func (interval Interval[T]) Lower() (T, bool) {
	return interval.lower, interval.lowerOpen
}

func (interval Interval[T]) Upper() (T, bool) {
	return interval.upper, interval.upperOpen
}

func (interval Interval[T]) Contains(value T) bool {
	lowerCmp := cmp.Compare[T](value, interval.lower)
	/* lowerCmp < 0:  Value is less than lower.
	 * lowerCmp == 0: Value is the same as lower. */
	if lowerCmp < 0 ||
		(interval.lowerOpen && lowerCmp == 0) {
		return false
	}

	upperCmp := cmp.Compare[T](value, interval.upper)
	/* upperCmp == 0: Value is the same as upper.
	 * upperCmp > 0:  Value is greater than upper.*/
	if (interval.upperOpen && upperCmp == 0) ||
		upperCmp > 0 {
		return false
	}

	return true
}

func (interval Interval[T]) Extends(other Interval[T]) bool {
	if interval.Empty() || other.Empty() {
		return false
	}

	if interval.Intersects(other) {
		return true
	}

	// Case 1: [1, 2] (2, 3]
	// Case 2: [2, 3] [1, 2)
	// Case 3: (2, 3] [1, 2]
	// Case 4: [1, 2) [2, 3]
	return other.lowerOpen && !interval.upperOpen && other.lower == interval.upper ||
		other.upperOpen && !interval.lowerOpen && other.upper == interval.lower ||
		interval.lowerOpen && !other.upperOpen && interval.lower == other.upper ||
		interval.upperOpen && !other.lowerOpen && interval.upper == other.lower
}

func (interval Interval[T]) Intersects(other Interval[T]) bool {
	if interval.Empty() || other.Empty() {
		return false
	}

	if other.lowerOpen && other.upperOpen &&
		interval.lowerOpen && interval.upperOpen {
		return interval.lower < other.upper && interval.upper > other.lower
	}

	return !other.lowerOpen && interval.Contains(other.lower) ||
		!other.upperOpen && interval.Contains(other.upper) ||
		!interval.upperOpen && other.Contains(interval.upper) ||
		!interval.lowerOpen && other.Contains(interval.lower)
}

func (interval *Interval[T]) Intersection(other Interval[T]) (Interval[T], error) {
	if !interval.Intersects(other) {
		return Interval[T]{}, ErrNoIntersection
	}

	intersectLower := interval.lower
	intersectUpper := interval.upper
	intersectLowerOpen := interval.lowerOpen
	intersectUpperOpen := interval.upperOpen

	if other.lower > interval.lower {
		intersectLower = other.lower
		intersectLowerOpen = other.lowerOpen
	}

	if other.upper < interval.upper {
		intersectUpper = other.upper
		intersectUpperOpen = other.upperOpen
	}

	return Interval[T]{
		lowerOpen: intersectLowerOpen,
		upperOpen: intersectUpperOpen,
		lower:     intersectLower,
		upper:     intersectUpper,
	}, nil
}

func (interval Interval[T]) Union(other Interval[T]) (Interval[T], error) {
	if !interval.Extends(other) {
		return Interval[T]{}, ErrNotExtendable
	}

	intersectLower := interval.lower
	intersectUpper := interval.upper
	intersectLowerOpen := interval.lowerOpen
	intersectUpperOpen := interval.upperOpen

	if other.lower < interval.lower {
		intersectLower = other.lower
		intersectLowerOpen = other.lowerOpen
	}

	if other.upper > interval.upper {
		intersectUpper = other.upper
		intersectUpperOpen = other.upperOpen
	}

	return Interval[T]{
		lowerOpen: intersectLowerOpen,
		upperOpen: intersectUpperOpen,
		lower:     intersectLower,
		upper:     intersectUpper,
	}, nil
}

type Intervals[T cmp.Ordered] []Interval[T]

func (intervals Intervals[T]) Len() int {
	return len(intervals)
}

func (intervals Intervals[T]) Less(i, j int) bool {
	return intervals[i].lower < intervals[j].lower
}

func (intervals Intervals[T]) Swap(i, j int) {
	(intervals)[i], (intervals)[j] = (intervals)[j], (intervals)[i]
}

func (intervals Intervals[T]) Contains(value T) bool {
	for idx := range intervals {
		if (intervals)[idx].Contains(value) {
			return true
		}
	}

	return false
}

func (intervals *Intervals[T]) Reduce() error {
	if !sort.IsSorted(intervals) {
		return ErrNotSorted
	}

	current := 0
	for i := 1; i < len(*intervals); i++ {
		if union, err := (*intervals)[current].Union((*intervals)[i]); err == nil {
			(*intervals)[current] = union
		} else {
			current++
			(*intervals)[current] = (*intervals)[i]
		}
	}

	// Truncate the slice to remove any remaining intervals
	*intervals = (*intervals)[:current+1]

	return nil
}
