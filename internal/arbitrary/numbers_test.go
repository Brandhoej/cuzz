package arbitrary

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestUnsignedLessThanWindow(t *testing.T) {
	var numbers map[uint8]struct{} = make(map[uint8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max uint8 = 0, 10
	var size int = int(max)

	bounds := fmt.Sprintf("[%d, %d)", min, max)

	in := func(val uint8) bool {
		return val >= min && val < max
	}

	for len(numbers) < size {
		sample := UnsignedLessThan[uint8](rng, max)

		if !in(sample) {
			t.Error("UnsignedLessThan[uint8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestUnsignedLessThanOrEqualWindow(t *testing.T) {
	var numbers map[uint8]struct{} = make(map[uint8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max uint8 = 0, 10
	var size int = int(max) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val uint8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := UnsignedLessThanOrEqual[uint8](rng, max)

		if !in(sample) {
			t.Error("UnsignedLessThanOrEqual[uint8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestUnsignedGreaterThanWindow(t *testing.T) {
	var numbers map[uint8]struct{} = make(map[uint8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max uint8 = math.MaxUint8 - 10, math.MaxUint8
	var size int = int(max - min)

	bounds := fmt.Sprintf("(%d, %d]", min, max)

	in := func(val uint8) bool {
		return val > min && val <= max
	}

	for len(numbers) < size {
		sample := UnsignedGreaterThan[uint8](rng, min)

		if !in(sample) {
			t.Error("UnsignedGreaterThan[uint8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestUnsignedGreaterThanOrEqualWindow(t *testing.T) {
	var numbers map[uint8]struct{} = make(map[uint8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max uint8 = math.MaxUint8 - 10, math.MaxUint8
	var size int = int(max-min) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val uint8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := UnsignedGreaterThanOrEqual[uint8](rng, min)

		if !in(sample) {
			t.Error("UnsignedGreaterThanOrEqual[uint8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestUnsignedInRangeWindow(t *testing.T) {
	var numbers map[uint8]struct{} = make(map[uint8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max uint8 = math.MaxUint8 - 10, math.MaxUint8
	var size int = int(max-min) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val uint8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := UnsignedInRange[uint8](rng, min, max)

		if !in(sample) {
			t.Error("UnsignedInRange[uint8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestSignedLessThanWindow(t *testing.T) {
	var numbers map[int8]struct{} = make(map[int8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max int8 = 0, math.MaxInt8 - 10
	var size int = int(max - min)

	bounds := fmt.Sprintf("[%d, %d)", min, max)

	in := func(val int8) bool {
		return val >= min && val < max
	}

	for len(numbers) < size {
		sample := SignedLessThan[int8](rng, max)

		if !in(sample) {
			t.Error("SignedLessThan[int8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestSignedLessThanOrEqualWindow(t *testing.T) {
	var numbers map[int8]struct{} = make(map[int8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max int8 = 0, math.MaxInt8 - 10
	var size int = int(max-min) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val int8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := SignedLessThanOrEqual[int8](rng, max)

		if !in(sample) {
			t.Error("SignedLessThanOrEqual[int8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestSignedGreaterThanWindow(t *testing.T) {
	var numbers map[int8]struct{} = make(map[int8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max int8 = math.MaxInt8 - 10, math.MaxInt8
	var size int = int(max - min)

	bounds := fmt.Sprintf("[%d, %d)", min, max)

	in := func(val int8) bool {
		return val >= min && val < max
	}

	for len(numbers) < size {
		sample := SignedGreaterThan[int8](rng, min)

		if !in(sample) {
			t.Error("SignedGreaterThan[int8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestSignedGreaterThanOrEqualWindow(t *testing.T) {
	var numbers map[int8]struct{} = make(map[int8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max int8 = math.MaxInt8 - 10, math.MaxInt8
	var size int = int(max-min) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val int8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := SignedGreaterThanOrEqual[int8](rng, min)

		if !in(sample) {
			t.Error("SignedGreaterThanOrEqual[int8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}

func TestSignedInRangeWindow(t *testing.T) {
	var numbers map[int8]struct{} = make(map[int8]struct{})
	rng := rand.New(rand.NewSource(1))
	var min, max int8 = -10, 10
	var size int = int(max-min) + 1

	bounds := fmt.Sprintf("[%d, %d]", min, max)

	in := func(val int8) bool {
		return val >= min && val <= max
	}

	for len(numbers) < size {
		sample := SignedInRange[int8](rng, min, max)

		if !in(sample) {
			t.Error("SignedInRange[int8] actual", sample, "expected inside", bounds)
		}

		if _, exists := numbers[sample]; !exists {
			numbers[sample] = struct{}{}
		}
	}
}
