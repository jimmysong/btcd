// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2013-2014 Dave Collins
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcec

// constant_time.go provides constant time implementations of useful mathematical operations. In addition, these
// functions return integers, using 0 or 1 to represent false or true respectively,
// which is useful for writing logic in terms of bitwise operators

// References
// These functions are based on the sample implementation in golang.org/src/crypto/subtle/constant_time.go
// Here we have refactored these functions for direct use in uint32 arithmetic

// ConstantTimeSelect returns x if v is 1 and y if v is 0.
// Its behavior is undefined if v takes any other value.
func SelectUint32(v, x, y uint32) uint32 { return ^(v-1)&x | (v-1)&y }

// ConstantTimeLessThan returns 1 if x < y and 0 otherwise.
// It works by checking the most significant bit, and then testing the rest of the bits by casting to int32
func LessThanUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32((diff >> 63) & 1)
}

func LessOrEqUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32(((diff - 1) >> 63) & 1)
}

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
func EqUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32((((diff - 1) ^ diff) >> 63) & 1)
}
