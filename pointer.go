// Copyright (C) 2024 neocotic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package pointers provides simple functions for working with pointers.
package pointers

// Bool returns a pointer to the given bool value.
func Bool(val bool) *bool {
	return &val
}

// Byte returns a pointer to the given byte value.
func Byte(val byte) *byte {
	return &val
}

// Complex64 returns a pointer to the given complex64 value.
func Complex64(val complex64) *complex64 {
	return &val
}

// Complex128 returns a pointer to the given complex128 value.
func Complex128(val complex128) *complex128 {
	return &val
}

// EmptyString is a convenient alias for ZeroString that returns a pointer to an empty string.
func EmptyString() *string {
	var val string
	return &val
}

// False is a convenient alias for ZeroBool that returns a pointer to a false bool value.
func False() *bool {
	var val bool
	return &val
}

// Float32 returns a pointer to the given float32 value.
func Float32(val float32) *float32 {
	return &val
}

// Float64 returns a pointer to the given float64 value.
func Float64(val float64) *float64 {
	return &val
}

// Int returns a pointer to the given int value.
func Int(val int) *int {
	return &val
}

// Int8 returns a pointer to the given int8 value.
func Int8(val int8) *int8 {
	return &val
}

// Int16 returns a pointer to the given int16 value.
func Int16(val int16) *int16 {
	return &val
}

// Int32 returns a pointer to the given int32 value.
func Int32(val int32) *int32 {
	return &val
}

// Int64 returns a pointer to the given int64 value.
func Int64(val int64) *int64 {
	return &val
}

// Rune returns a pointer to the given rune value.
func Rune(val rune) *rune {
	return &val
}

// String returns a pointer to the given string value.
func String(val string) *string {
	return &val
}

// True is a convenient alias for Bool(true) that returns a pointer to a true bool value.
func True() *bool {
	val := true
	return &val
}

// Uint returns a pointer to the given uint value.
func Uint(val uint) *uint {
	return &val
}

// Uint8 returns a pointer to the given uint8 value.
func Uint8(val uint8) *uint8 {
	return &val
}

// Uint16 returns a pointer to the given uint16 value.
func Uint16(val uint16) *uint16 {
	return &val
}

// Uint32 returns a pointer to the given uint32 value.
func Uint32(val uint32) *uint32 {
	return &val
}

// Uint64 returns a pointer to the given uint64 value.
func Uint64(val uint64) *uint64 {
	return &val
}

// Uintptr returns a pointer to the given uintptr value.
func Uintptr(val uintptr) *uintptr {
	return &val
}

// Value returns a pointer to the given value.
func Value[T any](val T) *T {
	return &val
}

// Zero returns a pointer to a zero value for T.
func Zero[T any]() *T {
	var val T
	return &val
}

// ZeroBool returns a pointer to a zero (false) bool value.
func ZeroBool() *bool {
	var val bool
	return &val
}

// ZeroByte returns a pointer to a zero byte value.
func ZeroByte() *byte {
	var val byte
	return &val
}

// ZeroComplex64 returns a pointer to a zero complex64 value.
func ZeroComplex64() *complex64 {
	var val complex64
	return &val
}

// ZeroComplex128 returns a pointer to a zero complex128 value.
func ZeroComplex128() *complex128 {
	var val complex128
	return &val
}

// ZeroFloat32 returns a pointer to a zero float32 value.
func ZeroFloat32() *float32 {
	var val float32
	return &val
}

// ZeroFloat64 returns a pointer to a zero float64 value.
func ZeroFloat64() *float64 {
	var val float64
	return &val
}

// ZeroInt returns a pointer to a zero int value.
func ZeroInt() *int {
	var val int
	return &val
}

// ZeroInt8 returns a pointer to a zero int8 value.
func ZeroInt8() *int8 {
	var val int8
	return &val
}

// ZeroInt16 returns a pointer to a zero int16 value.
func ZeroInt16() *int16 {
	var val int16
	return &val
}

// ZeroInt32 returns a pointer to a zero int32 value.
func ZeroInt32() *int32 {
	var val int32
	return &val
}

// ZeroInt64 returns a pointer to a zero int64 value.
func ZeroInt64() *int64 {
	var val int64
	return &val
}

// ZeroRune returns a pointer to a zero rune value.
func ZeroRune() *rune {
	var val rune
	return &val
}

// ZeroString returns a pointer to a zero (empty) string value.
func ZeroString() *string {
	var val string
	return &val
}

// ZeroUint returns a pointer to a zero uint value.
func ZeroUint() *uint {
	var val uint
	return &val
}

// ZeroUint8 returns a pointer to a zero uint8 value.
func ZeroUint8() *uint8 {
	var val uint8
	return &val
}

// ZeroUint16 returns a pointer to a zero uint16 value.
func ZeroUint16() *uint16 {
	var val uint16
	return &val
}

// ZeroUint32 returns a pointer to a zero uint32 value.
func ZeroUint32() *uint32 {
	var val uint32
	return &val
}

// ZeroUint64 returns a pointer to a zero uint64 value.
func ZeroUint64() *uint64 {
	var val uint64
	return &val
}

// ZeroUintptr returns a pointer to a zero uintptr value.
func ZeroUintptr() *uintptr {
	var val uintptr
	return &val
}
