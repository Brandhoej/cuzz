package arbitrary

import (
	"math"

	"golang.org/x/exp/constraints"
)

const (
	MinFloat32 = -0x1p127
	MinFloat64 = -0x1p1023

	MinUint8  = uint8(0)
	MinUint16 = uint16(0)
	MinUint32 = uint32(0)
	MinUint64 = uint64(0)
	MinUint   = uint(0)

	MaxUint8  = ^uint8(0)
	MaxUint16 = ^uint16(0)
	MaxUint32 = ^uint32(0)
	MaxUint64 = ^uint64(0)
	MaxUint   = ^uint(0)
)

func MinOf[T constraints.Integer | constraints.Float]() (value T) {
	switch any(value).(type) {
	case int8:
		var min int8 = math.MinInt8
		value = T(min)
	case int16:
		var min int16 = math.MinInt16
		value = T(min)
	case int32:
		var min int32 = math.MinInt32
		value = T(min)
	case int64:
		var min int64 = math.MinInt64
		value = T(min)
	case int:
		var min int = math.MinInt
		value = T(min)
	case uint8:
		var min uint8 = MinUint8
		value = T(min)
	case uint16:
		var min uint16 = MinUint16
		value = T(min)
	case uint32:
		var min uint32 = MinUint32
		value = T(min)
	case uint64:
		var min uint64 = MinUint64
		value = T(min)
	case uint:
		var min uint = MinUint
		value = T(min)
	case float32:
		var v float32 = MinFloat32
		value = T(v)
	case float64:
		var v float64 = MinFloat64
		value = T(v)
	default:
		panic("invalid type argument to MinOf")
	}
	return
}

func MaxOf[T constraints.Integer | constraints.Float]() (value T) {
	switch any(value).(type) {
	case int8:
		var max int8 = math.MaxInt8
		value = T(max)
	case int16:
		var max int16 = math.MaxInt16
		value = T(max)
	case int32:
		var max int32 = math.MaxInt32
		value = T(max)
	case int64:
		var max int64 = math.MaxInt64
		value = T(max)
	case int:
		var max int = math.MaxInt
		value = T(max)
	case uint8:
		var max uint8 = MaxUint8
		value = T(max)
	case uint16:
		var max uint16 = MaxUint16
		value = T(max)
	case uint32:
		var max uint32 = MaxUint32
		value = T(max)
	case uint64:
		var max uint64 = MaxUint64
		value = T(max)
	case uint:
		var max uint = MaxUint
		value = T(max)
	case float32:
		var v float32 = math.MaxFloat32
		value = T(v)
	case float64:
		var v float64 = math.MaxFloat64
		value = T(v)
	default:
		panic("invalid type argument to MaxOf")
	}
	return
}
