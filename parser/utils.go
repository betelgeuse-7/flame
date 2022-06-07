package parser

import (
	"flame/token"
	"math"
	"strconv"
)

// ? These functions look very alike. vvvv

func checkIsUint(dt token.TokenType, val string) bool {
	if dt == token.T_Uint32Kw {
		var maxU32Val uint32 = math.MaxUint32
		u32Val, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return false
		}
		if u32Val > uint64(maxU32Val) {
			return false
		}
		return true
	}
	// dt is token.T_UintKw
	var maxU64Val uint64 = math.MaxUint64
	u64Val, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return false
	}
	if u64Val > maxU64Val {
		return false
	}
	return true
}

func checkIsInt(dt token.TokenType, val string) bool {
	if dt == token.T_Int32Kw {
		var maxI32Val int32 = math.MaxInt32
		i32Val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return false
		}
		if i32Val > int64(maxI32Val) {
			return false
		}
		return true
	}
	var maxI64Val int64 = math.MaxInt64
	i64Val, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false
	}
	if i64Val > maxI64Val {
		return false
	}
	return true
}

func checkIsFloat(dt token.TokenType, val string) bool {
	if dt == token.T_Float32Kw {
		var maxF32Val float32 = math.MaxFloat32
		f32Val, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return false
		}
		if f32Val > float64(maxF32Val) {
			return false
		}
		return true
	}
	var maxF64Val float64 = math.MaxFloat64
	f64Val, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false
	}
	if f64Val > maxF64Val {
		return false
	}
	return true
}
