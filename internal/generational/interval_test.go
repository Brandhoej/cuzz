package generational

import (
	"testing"
)

func TestIntersects(t *testing.T) {
	tests := []struct {
		name     string
		lhs      Interval[int]
		rhs      Interval[int]
		expected bool
	}{
		{
			name:     "Intervals with overlap",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 3, upper: 8},
			expected: true,
		},
		{
			name:     "Intervals with no overlap",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 6, upper: 10},
			expected: false,
		},
		{
			name:     "Intervals with shared boundary",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 5, upper: 10},
			expected: true,
		},
		{
			name:     "Intervals with shared closed boundary",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 5, upper: 10},
			expected: true,
		},
		{
			name:     "Intervals with shared open boundary",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 5, upper: 10},
			expected: false,
		},
		{
			name:     "Intervals with the same closed boundary",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			expected: true,
		},
		{
			name:     "Intervals with the same open boundary",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			expected: true,
		},
		{
			name:     "Intervals with a single point overlap",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 5, upper: 10},
			expected: true,
		},
		{
			name:     "Empty interval with non-empty interval",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 5, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 10},
			expected: false,
		},
		{
			name:     "Empty interval with empty interval",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 1},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 1},
			expected: false,
		},
		{
			name:     "Overlapping intervals with the same upper bound",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 3, upper: 5},
			expected: true,
		},
		{
			name:     "Overlapping intervals with the same lower bound",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 3},
			expected: true,
		},
		{
			name:     "Overlapping intervals with the same bounds",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			expected: true,
		},
		{
			name:     "Overlapping intervals with open and closed bounds",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 4, upper: 8},
			expected: true,
		},
		{
			name:     "Overlapping intervals with the same open bounds",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 3, upper: 5},
			expected: true,
		},
		{
			name:     "Overlapping intervals with the same closed bounds",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 5},
			expected: true,
		},
		{
			name:     "Disjoint intervals with the same open and closed bounds",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 3, upper: 5},
			expected: false,
		},
		{
			name:     "Disjoint intervals with open bounds and shared closed boundary",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: false, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 5, upper: 10},
			expected: false,
		},
		{
			name:     "Disjoint intervals with closed bounds and shared open boundary",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 5, upper: 10},
			expected: false,
		},
		{
			name:     "Overlapping intervals with shared open and closed boundaries",
			lhs:      Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 5},
			expected: true,
		},
		{
			name:     "Overlapping intervals with the same open and closed boundaries",
			lhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 5},
			rhs:      Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 5},
			expected: true,
		},
	}

	for _, test := range tests {
		intersects := test.lhs.Intersects(test.rhs)
		if intersects != test.expected {
			t.Error(test.name, "got", intersects, "but expected", test.expected)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name         string
		lhs          Interval[int]
		rhs          Interval[int]
		intersection Interval[int]
		err          error
	}{
		{
			// [1, 5] ∩ [3, 8] -> [3, 5]
			name:         "Overlapping closed intervals",
			lhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 8},
			intersection: Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 5},
			err:          nil,
		},
		{
			// (1, 5) ∩ (3, 8) -> (3, 5)
			name:         "Overlapping open intervals",
			lhs:          Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:          Interval[int]{lowerOpen: true, upperOpen: true, lower: 3, upper: 8},
			intersection: Interval[int]{lowerOpen: true, upperOpen: true, lower: 3, upper: 5},
			err:          nil,
		},
		{
			// [1, 5] ∩ (3, 8) -> (3, 5]
			name:         "Overlapping closed and open intervals",
			lhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 5},
			rhs:          Interval[int]{lowerOpen: true, upperOpen: true, lower: 3, upper: 8},
			intersection: Interval[int]{lowerOpen: true, upperOpen: false, lower: 3, upper: 5},
			err:          nil,
		},
		{
			// (1, 5) ∩ [3, 8) -> [3, 5)
			name:         "Overlapping open and half-open intervals",
			lhs:          Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 5},
			rhs:          Interval[int]{lowerOpen: false, upperOpen: true, lower: 3, upper: 8},
			intersection: Interval[int]{lowerOpen: false, upperOpen: true, lower: 3, upper: 5},
			err:          nil,
		},
		{
			// [1, 1] ∩ [2, 2] -> ErrNoIntersection
			name:         "Disjoint single point closed intervals",
			lhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 1},
			rhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 2},
			intersection: Interval[int]{},
			err:          ErrNoIntersection,
		},
		{
			// (1, 1) ∩ [2, 2] -> ErrNoIntersection
			name:         "Disjoint open empty and closed single point intervals",
			lhs:          Interval[int]{lowerOpen: true, upperOpen: true, lower: 1, upper: 1},
			rhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 2},
			intersection: Interval[int]{},
			err:          ErrNoIntersection,
		},
		{
			// [1, 1] ∩ [1, 2] -> [1, 1]
			name:         "Overlapping single point closed and closed intervals",
			lhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 1},
			rhs:          Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			intersection: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 1},
			err:          nil,
		},
	}

	for _, test := range tests {
		intersection, err := test.lhs.Intersection(test.rhs)
		if intersection != test.intersection {
			t.Error(test.name, "intersection was", intersection, "but expected", test.intersection)
		}
		if err != test.err {
			t.Error(test.name, "error was", err, "but expected", test.err)
		}
	}
}

func TestExtends(t *testing.T) {
	tests := []struct {
		name    string
		lhs     Interval[int]
		rhs     Interval[int]
		extends bool
	}{
		{
			// [1, 2] (2, 3] -> true
			name:    "Extended by closed upper and open lower",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:     Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			extends: true,
		},
		{
			// [2, 3] [1, 2) -> true
			name:    "Extended by closed lower and open upper",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 2},
			extends: true,
		},
		{
			// (2, 3] [1, 2] -> true
			name:    "Extended by closed upper and open lower",
			lhs:     Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			extends: true,
		},
		{
			// [1, 2) [2, 3] -> true
			name:    "Extended by closed lower and open upper",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 2},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			extends: true,
		},
		{
			// [1, 2] [3, 4] -> false
			name:    "Not extended by closed intervals",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 4},
			extends: false,
		},
		{
			// [3, 4] [1, 2] -> false
			name:    "Not extended by closed intervals",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 4},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			extends: false,
		},
		{
			// [1, 2] [2, 3] -> true
			name:    "Extended by closed intervals with common bound",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			extends: true,
		},
		{
			// [2, 3] [1, 2] -> true
			name:    "Extended by closed intervals with common bound",
			lhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			rhs:     Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			extends: true,
		},
	}

	for _, test := range tests {
		extends := test.lhs.Extends(test.rhs)
		if extends != test.extends {
			t.Error(test.name, "was", extends, "but expected", test.extends)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name  string
		lhs   Interval[int]
		rhs   Interval[int]
		union Interval[int]
		err   error
	}{
		{
			// [1, 2] (2, 3] -> true
			name:  "Extended by closed upper and open lower",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:   Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
		{
			// [2, 3] [1, 2) -> true
			name:  "Extended by closed lower and open upper",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 2},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
		{
			// (2, 3] [1, 2] -> true
			name:  "Extended by closed upper and open lower",
			lhs:   Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
		{
			// [1, 2) [2, 3] -> true
			name:  "Extended by closed lower and open upper",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 2},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
		{
			// [1, 2] [3, 4] -> false
			name:  "Not extended by closed intervals",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 4},
			union: Interval[int]{},
			err:   ErrNotExtendable,
		},
		{
			// [3, 4] [1, 2] -> false
			name:  "Not extended by closed intervals",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 3, upper: 4},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			union: Interval[int]{},
			err:   ErrNotExtendable,
		},
		{
			// [1, 2] [2, 3] -> true
			name:  "Extended by closed intervals with common bound",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
		{
			// [2, 3] [1, 2] -> true
			name:  "Extended by closed intervals with common bound",
			lhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			rhs:   Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
			union: Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			err:   nil,
		},
	}

	for _, test := range tests {
		union, err := test.lhs.Union(test.rhs)
		if union != test.union {
			t.Error(test.name, "union was", union, "but expected", test.union)
		}
		if err != test.err {
			t.Error(test.name, "error was", union, "but expected", test.union)
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		original Intervals[int]
		reduced  Intervals[int]
		err      error
	}{
		{
			name: "One interval is already reduced",
			original: Intervals[int]{
				{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			},
			err: nil,
		},
		{
			name: "Extended by closed upper and open lower",
			original: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
				Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			},
			err: nil,
		},
		{
			name: "Extended by closed upper and open lower",
			original: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
				Interval[int]{lowerOpen: true, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			},
			err: nil,
		},
		{
			name: "Extended by closed upper and open lower",
			original: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: true, lower: 1, upper: 2},
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			},
			err: nil,
		},
		{
			name: "Extended by closed intervals with common bound",
			original: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			},
			err: nil,
		},
		{
			name: "Extended by closed intervals with common bound",
			original: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 2},
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 2, upper: 3},
			},
			reduced: Intervals[int]{
				Interval[int]{lowerOpen: false, upperOpen: false, lower: 1, upper: 3},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		test.original.Reduce()

		if len(test.original) != len(test.reduced) {
			t.Error(test.name, "length of the reduction was", len(test.reduced), "but expected", len(test.original))
		}

		minLength := len(test.original)
		if minLength > len(test.reduced) {
			minLength = len(test.reduced)
		}

		for i := 0; i < minLength; i++ {
			if test.original[i] != test.reduced[i] {
				t.Error(test.name, "interval at", i, "was", test.original[i], "but expected", test.reduced[i])
			}
		}
	}
}
