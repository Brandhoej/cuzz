package arbitrary

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// UnsignedLessThan returns a pseudo-random T-value in [0, n).
func UnsignedLessThan[T constraints.Unsigned](rng *rand.Rand, n T) T {
	if n == T(0) {
		panic("invalid argument to UnsignedLessThan")
	}

	// 0b_0100_0000 gets mask 0b_0011_1111 and all values after the mask has been applied will be less than n.
	if (n & (n - 1)) == 0 {
		return T(rng.Uint64()) & (n - 1)
	}

	// To remove bias we have to find the larges number which n is divisible with.
	maxRandomValue := ^T(0) - (^T(0) % n)

	randomValue := T(rng.Uint64())
	for randomValue > maxRandomValue {
		randomValue = T(rng.Uint64())
	}

	return randomValue % n
}

// UnsignedLessThanOrEqual returns a pseudo-random T-value in [0, n].
func UnsignedLessThanOrEqual[T constraints.Unsigned](rng *rand.Rand, n T) T {
	if n == ^T(0) {
		return T(rng.Uint64())
	}

	return UnsignedInRange[T](rng, 0, n)
}

// UnsignedGreaterThan returns a pseudo-random T-value in (n, ^T(0)].
func UnsignedGreaterThan[T constraints.Unsigned](rng *rand.Rand, n T) T {
	if n == ^T(0) {
		panic("invalid argument to UnsignedGreaterThan")
	}

	if n == ^T(0)-1 {
		return ^T(0)
	}

	return UnsignedInRange[T](rng, n+1, ^T(0))
}

// UnsignedGreaterThan returns a pseudo-random T-value in [n, ^T(0)].
func UnsignedGreaterThanOrEqual[T constraints.Unsigned](rng *rand.Rand, n T) T {
	if n == ^T(0) {
		return n
	}

	return UnsignedInRange[T](rng, n, ^T(0))
}

// UnsignedInRange returns a pseudo-random T-value in [min, max].
func UnsignedInRange[T constraints.Unsigned](rng *rand.Rand, min, max T) T {
	if min > max {
		panic("invalid argument to UnsignedInRange")
	}

	if min == max {
		return min
	}

	return min + UnsignedLessThan[T](rng, max-min+1)
}

// SignedLessThan returns a pseudo-random T-value in [0, n).
func SignedLessThan[T constraints.Signed](rng *rand.Rand, n T) T {
	if n == 0 {
		panic("invalid argument to SignedLessThan")
	}

	return SignedInRange[T](rng, 0, n-1)
}

// SignedLessThanOrEqual returns a pseudo-random T-value in [0, n].
func SignedLessThanOrEqual[T constraints.Signed](rng *rand.Rand, n T) T {
	if n == ^T(0) {
		return T(rng.Int63())
	}

	return SignedInRange[T](rng, 0, n)
}

// SignedGreaterThan returns a pseudo-random T-value in [n, max).
func SignedGreaterThan[T constraints.Signed](rng *rand.Rand, n T) T {
	max := MaxOf[T]()
	if n == max {
		panic("invalid argument to SignedGreaterThan")
	}

	return SignedInRange[T](rng, n, max-1)
}

// SignedGreaterThanOrEqual returns a pseudo-random T-value in [n, max].
func SignedGreaterThanOrEqual[T constraints.Signed](rng *rand.Rand, n T) T {
	max := MaxOf[T]()
	if n == max {
		return n
	}

	return SignedInRange[T](rng, n, max)
}

// SignedGreaterThanOrEqual returns a pseudo-random T-value in [min, max].
func SignedInRange[T constraints.Signed](rng *rand.Rand, min, max T) T {
	diff := safeUnsignedDifference[T, uint64](min, max)
	return safeUnsignedAddition[T, uint64](
		min, UnsignedLessThanOrEqual[uint64](rng, diff),
	)
}

// FloatLessThan returns a pseudo-random T-value in [0, n).
func FloatLessThan[T constraints.Float](rng *rand.Rand, n T) T {
	return T(float64(rng.Int63n(1<<53))/(1<<53)) * n
}

// FloatLessThanOrEqual returns a pseudo-random T-value in [0, n].
func FloatLessThanOrEqual[T constraints.Float](rng *rand.Rand, n T) T {
	return T(float64(rng.Int63())/(1<<63)) * n
}

// FloatGreaterThan returns a pseudo-random T-value in [n, max).
func FloatGreaterThan[T constraints.Float](rng *rand.Rand, n T) T {
	return lerp[T](n, MaxOf[T](), FloatLessThan[T](rng, 1.0))
}

// FloatGreaterThanOrEqual returns a pseudo-random T-value in [n, max].
func FloatGreaterThanOrEqual[T constraints.Float](rng *rand.Rand, n T) T {
	return lerp[T](n, MaxOf[T](), FloatLessThanOrEqual[T](rng, 1.0))
}

// FloatInRange returns a pseudo-random T-value in [min, max].
func FloatInRange[T constraints.Float](rng *rand.Rand, min, max T) T {
	return lerp[T](min, max, FloatLessThanOrEqual[T](rng, 1))
}

// LessThan returns a pseudo-random T-value in [0, n).
func LessThan[T constraints.Float | constraints.Integer](rng *rand.Rand, n T) (value T) {
	switch any(value).(type) {
	case int8:
		value = T(SignedLessThan[int8](rng, int8(n)))
	case int16:
		value = T(SignedLessThan[int16](rng, int16(n)))
	case int32:
		value = T(SignedLessThan[int32](rng, int32(n)))
	case int64:
		value = T(SignedLessThan[int64](rng, int64(n)))
	case int:
		value = T(SignedLessThan[int](rng, int(n)))
	case uint8:
		value = T(UnsignedLessThan[uint8](rng, uint8(n)))
	case uint16:
		value = T(UnsignedLessThan[uint16](rng, uint16(n)))
	case uint32:
		value = T(UnsignedLessThan[uint32](rng, uint32(n)))
	case uint64:
		value = T(UnsignedLessThan[uint64](rng, uint64(n)))
	case uint:
		value = T(UnsignedLessThan[uint](rng, uint(n)))
	case float32:
		value = T(FloatLessThan[float32](rng, float32(n)))
	case float64:
		value = T(FloatLessThan[float64](rng, float64(n)))
	}
	return
}

// LessThanOrEqual returns a pseudo-random T-value in [0, n].
func LessThanOrEqual[T constraints.Float | constraints.Integer](rng *rand.Rand, n T) (value T) {
	switch any(value).(type) {
	case int8:
		value = T(SignedLessThanOrEqual[int8](rng, int8(n)))
	case int16:
		value = T(SignedLessThanOrEqual[int16](rng, int16(n)))
	case int32:
		value = T(SignedLessThanOrEqual[int32](rng, int32(n)))
	case int64:
		value = T(SignedLessThanOrEqual[int64](rng, int64(n)))
	case int:
		value = T(SignedLessThanOrEqual[int](rng, int(n)))
	case uint8:
		value = T(UnsignedLessThanOrEqual[uint8](rng, uint8(n)))
	case uint16:
		value = T(UnsignedLessThanOrEqual[uint16](rng, uint16(n)))
	case uint32:
		value = T(UnsignedLessThanOrEqual[uint32](rng, uint32(n)))
	case uint64:
		value = T(UnsignedLessThanOrEqual[uint64](rng, uint64(n)))
	case uint:
		value = T(UnsignedLessThanOrEqual[uint](rng, uint(n)))
	case float32:
		value = T(FloatLessThanOrEqual[float32](rng, float32(n)))
	case float64:
		value = T(FloatLessThanOrEqual[float64](rng, float64(n)))
	}
	return
}

// GreaterThan returns a pseudo-random T-value in [n, max).
func GreaterThan[T constraints.Float | constraints.Integer](rng *rand.Rand, n T) (value T) {
	switch any(value).(type) {
	case int8:
		value = T(SignedGreaterThan[int8](rng, int8(n)))
	case int16:
		value = T(SignedGreaterThan[int16](rng, int16(n)))
	case int32:
		value = T(SignedGreaterThan[int32](rng, int32(n)))
	case int64:
		value = T(SignedGreaterThan[int64](rng, int64(n)))
	case int:
		value = T(SignedGreaterThan[int](rng, int(n)))
	case uint8:
		value = T(UnsignedGreaterThan[uint8](rng, uint8(n)))
	case uint16:
		value = T(UnsignedGreaterThan[uint16](rng, uint16(n)))
	case uint32:
		value = T(UnsignedGreaterThan[uint32](rng, uint32(n)))
	case uint64:
		value = T(UnsignedGreaterThan[uint64](rng, uint64(n)))
	case uint:
		value = T(UnsignedGreaterThan[uint](rng, uint(n)))
	case float32:
		value = T(FloatGreaterThan[float32](rng, float32(n)))
	case float64:
		value = T(FloatGreaterThan[float64](rng, float64(n)))
	}
	return
}

// GreaterThanOrEqual returns a pseudo-random T-value in [n, max].
func GreaterThanOrEqual[T constraints.Float | constraints.Integer](rng *rand.Rand, n T) (value T) {
	switch any(value).(type) {
	case int8:
		value = T(SignedGreaterThanOrEqual[int8](rng, int8(n)))
	case int16:
		value = T(SignedGreaterThanOrEqual[int16](rng, int16(n)))
	case int32:
		value = T(SignedGreaterThanOrEqual[int32](rng, int32(n)))
	case int64:
		value = T(SignedGreaterThanOrEqual[int64](rng, int64(n)))
	case int:
		value = T(SignedGreaterThanOrEqual[int](rng, int(n)))
	case uint8:
		value = T(UnsignedGreaterThanOrEqual[uint8](rng, uint8(n)))
	case uint16:
		value = T(UnsignedGreaterThanOrEqual[uint16](rng, uint16(n)))
	case uint32:
		value = T(UnsignedGreaterThanOrEqual[uint32](rng, uint32(n)))
	case uint64:
		value = T(UnsignedGreaterThanOrEqual[uint64](rng, uint64(n)))
	case uint:
		value = T(UnsignedGreaterThanOrEqual[uint](rng, uint(n)))
	case float32:
		value = T(FloatGreaterThanOrEqual[float32](rng, float32(n)))
	case float64:
		value = T(FloatGreaterThanOrEqual[float64](rng, float64(n)))
	}
	return
}

// InRange returns a pseudo-random T-value in [min, max].
func InRange[T constraints.Float | constraints.Integer](rng *rand.Rand, min, max T) (value T) {
	switch any(value).(type) {
	case int8:
		value = T(SignedInRange[int8](rng, int8(min), int8(max)))
	case int16:
		value = T(SignedInRange[int16](rng, int16(min), int16(max)))
	case int32:
		value = T(SignedInRange[int32](rng, int32(min), int32(max)))
	case int64:
		value = T(SignedInRange[int64](rng, int64(min), int64(max)))
	case int:
		value = T(SignedInRange[int](rng, int(min), int(max)))
	case uint8:
		value = T(UnsignedInRange[uint8](rng, uint8(min), uint8(max)))
	case uint16:
		value = T(UnsignedInRange[uint16](rng, uint16(min), uint16(max)))
	case uint32:
		value = T(UnsignedInRange[uint32](rng, uint32(min), uint32(max)))
	case uint64:
		value = T(UnsignedInRange[uint64](rng, uint64(min), uint64(max)))
	case uint:
		value = T(UnsignedInRange[uint](rng, uint(min), uint(max)))
	case float32:
		value = T(FloatInRange[float32](rng, float32(min), float32(max)))
	case float64:
		value = T(FloatInRange[float64](rng, float64(min), float64(max)))
	}
	return
}
