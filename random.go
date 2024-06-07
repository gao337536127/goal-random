package GoalRandom

import (
	"encoding/binary"
	"math"
)

func RandBytes(length int) []byte {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = <-byteCache
	}

	return bytes
}

func Sign() int {
	a := RandUint16()
	b := RandUint16()
	if a > b {
		return 1
	}
	return -1
}

func RandUInt64() uint64 {
	bytes := RandBytes(8)
	return binary.NativeEndian.Uint64(bytes)
}

func RandInt64() int64 {
	return int64(Sign()) * int64(RandUInt64())
}

func RandUint32() uint32 {
	bytes := RandBytes(4)
	return binary.NativeEndian.Uint32(bytes)
}

func RandInt32() int32 {
	return int32(Sign()) * int32(RandUint32())
}

func RandUint16() uint16 {
	bytes := RandBytes(2)
	return binary.NativeEndian.Uint16(bytes)
}

func RandInt16() int16 {
	return int16(Sign()) * int16(RandUint16())
}

func RandFloat64() float64 {
	X := float64(RandUInt64())
	Y := float64(RandUInt64())
	Z := X / Y
	_, div := math.Modf(Z)
	return float64(Sign()) * div
}

func RandInt32Slice(length int) []int32 {
	result := make([]int32, length)
	for i := 0; i < length; i++ {
		result[i] = RandInt32()
	}
	return result
}
