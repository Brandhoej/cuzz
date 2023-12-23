package arbitrary

import (
	"math"
	"testing"
)

func TestSafeUnsignedDifferenceInt8(t *testing.T) {
	for i := uint16(0); i <= math.MaxUint8; i++ {
		min := int8(math.MinInt8)
		max := int8(int16(math.MaxInt8) - int16(i))

		actualMinMax := safeUnsignedDifference[int8, uint8](min, max)
		actualMaxMin := safeUnsignedDifference[int8, uint8](max, min)
		expected := math.MaxUint8 - uint8(i)

		if actualMinMax != expected {
			t.Error("safeUnsignedDifference[int8, uint8](min, max) actual", actualMinMax, "expected", expected)
		}

		if actualMaxMin != expected {
			t.Error("safeUnsignedDifference[int8, uint8](max, min) actual", actualMinMax, "expected", expected)
		}
	}
}

func TestSafeUnsignedDifference(t *testing.T) {
	if actual := safeUnsignedDifference[int8, uint8](math.MinInt8, math.MaxInt8); actual != math.MaxUint8 {
		t.Error("safeUnsignedDifference[int8, uint8](min, max) actual", actual, "expected", math.MaxUint8)
	}
	if actual := safeUnsignedDifference[int16, uint16](math.MinInt16, math.MaxInt16); actual != math.MaxUint16 {
		t.Error("safeUnsignedDifference[int16, uint16](min, max) actual", actual, "expected", math.MaxUint16)
	}
	if actual := safeUnsignedDifference[int32, uint32](math.MinInt32, math.MaxInt32); actual != math.MaxUint32 {
		t.Error("safeUnsignedDifference[int32, uint32](min, max) actual", actual, "expected", math.MaxUint32)
	}
	if actual := safeUnsignedDifference[int64, uint64](math.MinInt64, math.MaxInt64); actual != math.MaxUint64 {
		t.Error("safeUnsignedDifference[int64, uint64](min, max) actual", actual, "expected", uint64(math.MaxUint64))
	}
	if actual := safeUnsignedDifference[int, uint](math.MinInt, math.MaxInt); actual != math.MaxUint {
		t.Error("safeUnsignedDifference[int64, uint64](min, max) actual", actual, "expected", math.MaxUint8)
	}
}

func TestSafeUnsignedAdditionInt8(t *testing.T) {
	for i := ^uint8(0); i > 0; i-- {
		actual := safeUnsignedAddition[int8, uint8](math.MinInt8, math.MaxUint8-i)
		expected := int8(math.MaxInt8 - i)

		if actual != expected {
			t.Error("safeUnsignedAddition[int8, uint8] actual", actual, "expected", expected)
		}
	}
}

func TestSafeUnsignedAddition(t *testing.T) {
	if actual := safeUnsignedAddition[int8, uint8](math.MinInt8, math.MaxUint8); actual != math.MaxInt8 {
		t.Error("safeUnsignedAddition[int8, uint8] actual", actual, "expected", math.MaxInt8)
	}
	if actual := safeUnsignedAddition[int16, uint16](math.MinInt16, math.MaxUint16); actual != math.MaxInt16 {
		t.Error("safeUnsignedAddition[int16, uint16] actual", actual, "expected", math.MaxInt16)
	}
	if actual := safeUnsignedAddition[int32, uint32](math.MinInt32, math.MaxUint32); actual != math.MaxInt32 {
		t.Error("safeUnsignedAddition[int32, uint32] actual", actual, "expected", math.MaxInt32)
	}
	if actual := safeUnsignedAddition[int64, uint64](math.MinInt64, math.MaxUint64); actual != math.MaxInt64 {
		t.Error("safeUnsignedAddition[int64, uint64] actual", actual, "expected", math.MaxInt64)
	}
	if actual := safeUnsignedAddition[int, uint](math.MinInt, math.MaxUint); actual != math.MaxInt {
		t.Error("safeUnsignedAddition[int, uint] actual", actual, "expected", math.MaxInt)
	}
}
