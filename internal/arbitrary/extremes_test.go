package arbitrary

import (
	"math"
	"testing"
)

func TestMinsAndMaxs(t *testing.T) {
	if actual := MinFloat32; actual != -0x1p127 {
		t.Error("MinFloat32 actual", actual, "expected", -0x1p127)
	}
	if actual := MinFloat64; actual != -0x1p1023 {
		t.Error("MinFloat32 actual", actual, "expected", -0x1p127)
	}
	if actual := MinUint8; actual != uint8(0) {
		t.Error("MinUint8 actual", actual, "expected", uint8(0))
	}
	if actual := MinUint16; actual != uint16(0) {
		t.Error("MinUint16 actual", actual, "expected", uint16(0))
	}
	if actual := MinUint32; actual != uint32(0) {
		t.Error("MinUint32 actual", actual, "expected", uint32(0))
	}
	if actual := MinUint64; actual != uint64(0) {
		t.Error("MinUint64 actual", actual, "expected", uint64(0))
	}
	if actual := MinUint; actual != uint(0) {
		t.Error("MinUint actual", actual, "expected", uint(0))
	}
	if actual := MaxUint8; actual != ^uint8(0) {
		t.Error("MaxUint8 actual", actual, "expected", ^uint8(0))
	}
	if actual := MaxUint16; actual != ^uint16(0) {
		t.Error("MaxUint16 actual", actual, "expected", ^uint16(0))
	}
	if actual := MaxUint32; actual != ^uint32(0) {
		t.Error("MaxUint32 actual", actual, "expected", ^uint32(0))
	}
	if actual := MaxUint64; actual != ^uint64(0) {
		t.Error("MaxUint64 actual", actual, "expected", ^uint64(0))
	}
	if actual := MaxUint; actual != ^uint(0) {
		t.Error("MaxUint actual", actual, "expected", ^uint(0))
	}
}

func TestMinOf(t *testing.T) {
	// Signed:
	if min := MinOf[int8](); min != math.MinInt8 {
		t.Error("MinOf int8 actual", min, "expected", math.MinInt8)
	}
	if min := MinOf[int16](); min != math.MinInt16 {
		t.Error("MinOf int16 actual", min, "expected", math.MinInt8)
	}
	if min := MinOf[int32](); min != math.MinInt32 {
		t.Error("MinOf int32 actual", min, "expected", math.MinInt8)
	}
	if min := MinOf[int64](); min != math.MinInt64 {
		t.Error("MinOf int64 actual", min, "expected", math.MinInt8)
	}
	if min := MinOf[int](); min != math.MinInt {
		t.Error("MinOf int actual", min, "expected", math.MinInt8)
	}

	// Unsigned:
	if min := MinOf[uint8](); min != MinUint8 {
		t.Error("MinOf uint8 actual", min, "expected", MinUint8)
	}
	if min := MinOf[uint16](); min != MinUint16 {
		t.Error("MinOf uint16 actual", min, "expected", MinUint16)
	}
	if min := MinOf[uint32](); min != MinUint32 {
		t.Error("MinOf uint32 actual", min, "expected", MinUint32)
	}
	if min := MinOf[uint64](); min != MinUint64 {
		t.Error("MinOf uint64 actual", min, "expected", MinUint64)
	}
	if min := MinOf[uint](); min != MinUint {
		t.Error("MinOf uint actual", min, "expected", MinUint)
	}

	// Float:
	if min := MinOf[float32](); min != -0x1p127 {
		t.Error("MinOf float32 actual", min, "expected", MinFloat32)
	}
	if min := MinOf[float64](); min != -0x1p1023 {
		t.Error("MinOf float64 actual", min, "expected", MinFloat64)
	}
}

func BenchmarkMinOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinOf[float64]()
	}
}

func TestMaxOf(t *testing.T) {
	// Signed:
	if max := MaxOf[int8](); max != math.MaxInt8 {
		t.Error("MaxOf int8 actual", max, "expected", math.MaxInt8)
	}
	if max := MaxOf[int16](); max != math.MaxInt16 {
		t.Error("MaxOf int16 actual", max, "expected", math.MaxInt8)
	}
	if max := MaxOf[int32](); max != math.MaxInt32 {
		t.Error("MaxOf int32 actual", max, "expected", math.MaxInt8)
	}
	if max := MaxOf[int64](); max != math.MaxInt64 {
		t.Error("MaxOf int64 actual", max, "expected", math.MaxInt8)
	}
	if max := MaxOf[int](); max != math.MaxInt {
		t.Error("MaxOf int actual", max, "expected", math.MaxInt8)
	}

	// Unsigned:
	if max := MaxOf[uint8](); max != MaxUint8 {
		t.Error("MaxOf uint8 actual", max, "expected", MaxUint8)
	}
	if max := MaxOf[uint16](); max != MaxUint16 {
		t.Error("MaxOf uint16 actual", max, "expected", MaxUint16)
	}
	if max := MaxOf[uint32](); max != MaxUint32 {
		t.Error("MaxOf uint32 actual", max, "expected", MaxUint32)
	}
	if max := MaxOf[uint64](); max != MaxUint64 {
		t.Error("MaxOf uint64 actual", max, "expected", MaxUint64)
	}
	if max := MaxOf[uint](); max != MaxUint {
		t.Error("MaxOf uint actual", max, "expected", MaxUint)
	}

	// Float:
	if max := MaxOf[float32](); max != math.MaxFloat32 {
		t.Error("MaxOf float32 actual", max, "expected", math.MaxFloat32)
	}
	if max := MaxOf[float64](); max != math.MaxFloat64 {
		t.Error("MaxOf float64 actual", max, "expected", math.MaxFloat32)
	}
}

func BenchmarkMaxOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxOf[float64]()
	}
}
