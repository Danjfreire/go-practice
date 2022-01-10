package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	Methods
	A method is a function with a special "receiver" argument.
	The receiver appear between the "func" keyword and the method name.
	Once declared, the method becomes available to the type referenced by the receiver.
*/

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Methods can be declared on non-struct types.
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
	You can create a pointer receiver if the method need to modify
	the receiver (like set methods on regular oop languages).
	Pointer receivers are also more efficient because there is no need to create a copy
	of the values to execute the method.
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())
}
