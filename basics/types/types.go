package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
	Go has the following basic types:

	string
	bool
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte (alias for uint8)
	rune (alias for int32)
	float32 float64
	complex64 complex128
*/

/*
	int, uint and uintptr are 32 or 64 bits depending on the operating system.
	The specific types that are unsigned or end with bits should only be used in specific cases
*/

/*
	Constants are declared with the const keyword and unlike variables, they cannot
	be declared using the := syntax
*/
const Pi = 3.14

func main() {
	var ToBe bool = false
	var MaxInt uint64 = 1<<64 - 1 // bit shift operation
	var z complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*
		Variables declared without a initial value are given a "zero value":
		0 for numeric types
		false for boolean
		"" for strings
	*/
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", s, s)

	/*
		Types can be converted using the expression : T(v) where T is the target type and
		v is the value to be converted
	*/
	var x, y int = 3, 4
	var w float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", w, w)
}
