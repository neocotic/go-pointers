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

package pointers

import (
	"fmt"
	"math"
	"testing"
)

func Test_Bool(t *testing.T) {
	testPointerFunc(t, Bool, false, true)
}

func Test_Byte(t *testing.T) {
	testPointerFunc(t, Byte, 0, 1, math.MaxUint8)
}

func Test_Complex64(t *testing.T) {
	testPointerFunc(t, Complex64, 0, complex(-1, 2), complex(3, 4))
}

func Test_Complex128(t *testing.T) {
	testPointerFunc(t, Complex128, 0, complex(-1, 2), complex(3, 4))
}

func Test_EmptyString(t *testing.T) {
	testPointerConstFunc(t, EmptyString, "")
}

func Test_False(t *testing.T) {
	testPointerConstFunc(t, False, false)
}

func Test_Float32(t *testing.T) {
	testPointerFunc(t, Float32, 0, -math.SmallestNonzeroFloat32, -math.MaxFloat32, math.SmallestNonzeroFloat32, math.MaxFloat32)
}

func Test_Float64(t *testing.T) {
	testPointerFunc(t, Float64, 0, -math.SmallestNonzeroFloat64, -math.MaxFloat64, math.SmallestNonzeroFloat64, math.MaxFloat64)
}

func Test_Int(t *testing.T) {
	testPointerFunc(t, Int, 0, -1, math.MinInt, 1, math.MaxInt)
}

func Test_Int8(t *testing.T) {
	testPointerFunc(t, Int8, 0, -1, math.MinInt8, 1, math.MaxInt8)
}

func Test_Int16(t *testing.T) {
	testPointerFunc(t, Int16, 0, -1, math.MinInt16, 1, math.MaxInt16)
}

func Test_Int32(t *testing.T) {
	testPointerFunc(t, Int32, 0, -1, math.MinInt32, 1, math.MaxInt32)
}

func Test_Int64(t *testing.T) {
	testPointerFunc(t, Int64, 0, -1, math.MinInt64, 1, math.MaxInt64)
}

func Test_Rune(t *testing.T) {
	testPointerFunc(t, Rune, 0, -1, math.MinInt32, 1, math.MaxInt32)
}

func Test_String(t *testing.T) {
	testPointerFunc(t, String, "", "foo", "bar")
}

func Test_True(t *testing.T) {
	testPointerConstFunc(t, True, true)
}

func Test_Uint(t *testing.T) {
	testPointerFunc(t, Uint, 0, 1, math.MaxUint)
}

func Test_Uint8(t *testing.T) {
	testPointerFunc(t, Uint8, 0, 1, math.MaxUint8)
}

func Test_Uint16(t *testing.T) {
	testPointerFunc(t, Uint16, 0, 1, math.MaxUint16)
}

func Test_Uint32(t *testing.T) {
	testPointerFunc(t, Uint32, 0, 1, math.MaxUint32)
}

func Test_Uint64(t *testing.T) {
	testPointerFunc(t, Uint64, 0, 1, math.MaxUint64)
}

func Test_Value(t *testing.T) {
	type test struct {
		String string
	}

	testPointerFunc(t, Value[test], test{}, test{"foo"}, test{"bar"})
}

func Test_Zero(t *testing.T) {
	type test struct {
		String string
	}

	testPointerConstFunc(t, Zero[test], test{})
}

func Test_ZeroBool(t *testing.T) {
	testPointerConstFunc(t, ZeroBool, false)
}

func Test_ZeroByte(t *testing.T) {
	testPointerConstFunc(t, ZeroByte, 0)
}

func Test_ZeroComplex64(t *testing.T) {
	testPointerConstFunc(t, ZeroComplex64, 0)
}

func Test_ZeroComplex128(t *testing.T) {
	testPointerConstFunc(t, ZeroComplex128, 0)
}

func Test_ZeroFloat32(t *testing.T) {
	testPointerConstFunc(t, ZeroFloat32, 0)
}

func Test_ZeroFloat64(t *testing.T) {
	testPointerConstFunc(t, ZeroFloat64, 0)
}

func Test_ZeroInt(t *testing.T) {
	testPointerConstFunc(t, ZeroInt, 0)
}

func Test_ZeroInt8(t *testing.T) {
	testPointerConstFunc(t, ZeroInt8, 0)
}

func Test_ZeroInt16(t *testing.T) {
	testPointerConstFunc(t, ZeroInt16, 0)
}

func Test_ZeroInt32(t *testing.T) {
	testPointerConstFunc(t, ZeroInt32, 0)
}

func Test_ZeroInt64(t *testing.T) {
	testPointerConstFunc(t, ZeroInt64, 0)
}

func Test_ZeroRune(t *testing.T) {
	testPointerConstFunc(t, ZeroRune, 0)
}

func Test_ZeroString(t *testing.T) {
	testPointerConstFunc(t, ZeroString, "")
}

func Test_ZeroUint(t *testing.T) {
	testPointerConstFunc(t, ZeroUint, 0)
}

func Test_ZeroUint8(t *testing.T) {
	testPointerConstFunc(t, ZeroUint8, 0)
}

func Test_ZeroUint16(t *testing.T) {
	testPointerConstFunc(t, ZeroUint16, 0)
}

func Test_ZeroUint32(t *testing.T) {
	testPointerConstFunc(t, ZeroUint32, 0)
}

func Test_ZeroUint64(t *testing.T) {
	testPointerConstFunc(t, ZeroUint64, 0)
}

func testPointerFunc[T comparable](t *testing.T, fn func(val T) *T, vals ...T) {
	for _, val := range vals {
		t.Run(fmt.Sprintf("with value: %v", val), func(t *testing.T) {
			ret := fn(val)
			if ret == nil {
				t.Errorf("unexpected pointer; want %v, got nil", val)
			} else if *ret != val {
				t.Errorf("unexpected pointer; want %v, got %v", val, *ret)
			}
		})
	}
}

func testPointerConstFunc[T comparable](t *testing.T, fn func() *T, expect T) {
	ret := fn()
	if ret == nil {
		t.Errorf("unexpected pointer; want %v, got nil", expect)
	} else if *ret != expect {
		t.Errorf("unexpected pointer; want %v, got %v", expect, *ret)
	}
}
