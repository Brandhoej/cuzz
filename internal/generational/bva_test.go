package generational

import (
	"math"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestBVA(t *testing.T) {
	tests := []struct {
		name  string
		first int
		last  int
		avg   int
		bva   []int
	}{
		{
			name:  "Negative and positive bounds",
			first: -10,
			last:  10,
			avg:   0,
			bva:   []int{-10, -9, 0, 9, 10},
		},
		{
			name:  "Single value interval",
			first: 0,
			last:  0,
			avg:   0,
			bva:   []int{0},
		},
		{
			name:  "Two value interval",
			first: 0,
			last:  1,
			avg:   0,
			bva:   []int{0, 1},
		},
		{
			name:  "Extreme values",
			first: math.MinInt,
			last:  math.MaxInt,
			avg:   0,
			bva:   []int{math.MinInt, math.MinInt + 1, 0, math.MaxInt - 1, math.MaxInt},
		},
	}

	for _, test := range tests {
		bva := BoundaryStrategyNumeric[int](
			test.first, test.last, 1,
			test.first, test.last,
			BVA[int],
		)

		if len(bva) != len(test.bva) {
			t.Error(test.name, "length of the BVA was", len(bva), "but expected", len(test.bva))
		}

		for _, value := range test.bva {
			if !slices.Contains[[]int, int](bva, value) {
				t.Error(test.name, "expected BVA to contain", value, "but it did not")
			}
		}

		for _, value := range bva {
			if !slices.Contains[[]int, int](test.bva, value) {
				t.Error(test.name, "did not expect BVA to contain", value, "but it did")
			}
		}
	}
}

func TestRA(t *testing.T) {
	tests := []struct {
		name  string
		first int
		last  int
		min   int
		max   int
		avg   int
		ra    []int
	}{
		{
			name:  "Negative and positive bounds",
			first: -10,
			last:  10,
			avg:   0,
			min:   -11,
			max:   11,
			ra:    []int{-11, -10, -9, 0, 9, 10, 11},
		},
		{
			name:  "Single value interval",
			first: 0,
			last:  0,
			avg:   0,
			min:   -1,
			max:   1,
			ra:    []int{-1, 0, 1},
		},
		{
			name:  "Two value interval",
			first: 0,
			last:  1,
			avg:   0,
			min:   -1,
			max:   2,
			ra:    []int{-1, 0, 1, 2},
		},
		{
			name:  "Extreme values",
			first: math.MinInt,
			last:  math.MaxInt,
			avg:   0,
			min:   math.MinInt,
			max:   math.MaxInt,
			ra:    []int{math.MinInt, math.MinInt + 1, 0, math.MaxInt - 1, math.MaxInt},
		},
		{
			name:  "Just around extreme values",
			first: math.MinInt + 1,
			last:  math.MaxInt - 1,
			avg:   0,
			min:   math.MinInt,
			max:   math.MaxInt,
			ra:    []int{math.MinInt, math.MinInt + 1, math.MinInt + 2, 0, math.MaxInt - 2, math.MaxInt - 1, math.MaxInt},
		},
	}

	for _, test := range tests {
		ra := BoundaryStrategyNumeric[int](
			test.first, test.last, 1,
			test.min, test.max,
			RA[int],
		)

		if len(ra) != len(test.ra) {
			t.Error(test.name, "length of the RA was", len(ra), "but expected", len(test.ra))
		}

		for _, value := range test.ra {
			if !slices.Contains[[]int, int](ra, value) {
				t.Error(test.name, "expected RA to contain", value, "but it did not")
			}
		}

		for _, value := range ra {
			if !slices.Contains[[]int, int](test.ra, value) {
				t.Error(test.name, "did not expect RA to contain", value, "but it did")
			}
		}
	}
}

func TestWCA(t *testing.T) {
	tests := []struct {
		name   string
		first1 int
		last1  int
		first2 int
		last2  int
		wca    [][]int
	}{
		{
			name: "WCA [1, 5] [2, 3]",
			// BVA1: {1, 2, 3, 4, 5}
			// BVA2: {2, 3}
			first1: 1,
			last1:  5,
			first2: 2,
			last2:  3,
			wca: [][]int{
				{1, 2},
				{1, 3},
				{2, 2},
				{2, 3},
				{3, 2},
				{3, 3},
				{4, 2},
				{4, 3},
				{5, 2},
				{5, 3},
			},
		},
		{
			name: "WCA [1, 1] [2, 3]",
			// BVA1: {1}
			// BVA2: {2, 3}
			first1: 1,
			last1:  1,
			first2: 2,
			last2:  3,
			wca: [][]int{
				{1, 2},
				{1, 3},
			},
		},
	}

	for _, test := range tests {
		bva1 := func() []int {
			return BoundaryStrategyNumeric[int](
				test.first1, test.last1, 1,
				test.first1, test.last1,
				BVA[int],
			)
		}
		bva2 := func() []int {
			return BoundaryStrategyNumeric[int](
				test.first2, test.last2, 1,
				test.first2, test.last2,
				BVA[int],
			)
		}
		wca := MBA[int]([]func() []int{bva1, bva2})

		if len(wca) <= 0 {
			t.Error(test.name, "WCA result should be larger than", len(wca))
		}

		for _, point := range wca {
			if len(point) != 2 {
				t.Error(test.name, "point length was", len(point), "expected", 2)
			}

			found := false
			for _, expected := range test.wca {
				if reflect.DeepEqual(point, expected) {
					found = true
					break
				}
			}

			if !found {
				t.Error(test.name, "point", point, "is not expected")
			}
		}
	}
}

func TestRWCA(t *testing.T) {
	tests := []struct {
		name   string
		first1 int
		last1  int
		first2 int
		last2  int
		wca    [][]int
	}{
		{
			name: "RWCA [1, 5] [2, 3]",
			// RA1: {0, 1, 2, 3, 4, 5, 6}
			// RA2: {1, 2, 3, 4}
			first1: 1,
			last1:  5,
			first2: 2,
			last2:  3,
			wca: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 1},
				{2, 2},
				{2, 3},
				{2, 4},
				{3, 1},
				{3, 2},
				{3, 3},
				{3, 4},
				{4, 1},
				{4, 2},
				{4, 3},
				{4, 4},
				{5, 1},
				{5, 2},
				{5, 3},
				{5, 4},
				{6, 1},
				{6, 2},
				{6, 3},
				{6, 4},
			},
		},
		{
			name: "RWCA [1, 1] [2, 3]",
			// RA1: {0, 1, 2}
			// RA2: {1, 2, 3, 4}
			first1: 1,
			last1:  1,
			first2: 2,
			last2:  3,
			wca: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 1},
				{2, 2},
				{2, 3},
				{2, 4},
			},
		},
	}

	for _, test := range tests {
		ra1 := func() []int {
			return BoundaryStrategyNumeric[int](
				test.first1, test.last1, 1,
				math.MinInt, math.MaxInt,
				RA[int],
			)
		}
		ra2 := func() []int {
			return BoundaryStrategyNumeric[int](
				test.first2, test.last2, 1,
				math.MinInt, math.MaxInt,
				RA[int],
			)
		}
		rwca := MBA[int]([]func() []int{ra1, ra2})

		if len(rwca) <= 0 {
			t.Error(test.name, "WCA result should be larger than", len(rwca))
		}

		for _, point := range rwca {
			if len(point) != 2 {
				t.Error(test.name, "point length was", len(point), "expected", 2)
			}

			found := false
			for _, expected := range test.wca {
				if reflect.DeepEqual(point, expected) {
					found = true
					break
				}
			}

			if !found {
				t.Error(test.name, "point", point, "is not expected")
			}
		}
	}
}
