package main

import (
	"fmt"
	"math"
)

/*
	Interfaces
	An interface type is defined as a set of method signatures
*/

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	var a Abser
	v := Vertex{3, 4}

	/*
		A value of interface can hold any value that implements it.
		a = v is possible because v which is a Vertex implements the Abs method.
		This would not be possible if the Abs method on Vertex used a pointer receiver.
	*/

	a = v

	fmt.Println(a)

	// ----------------------------------------------------------------------------------

	/*
		Type Assertions
		With a type assertion you get access to the interface underlying concrete value
	*/

	// empty interface
	var i interface{} = "hello"

	/*
		In this example "ok" is a boolean that checks if the underlying value
		is of type string, and "s" stores the underlying value if it is true.
		If its false the value will be the zero value of the type.
	*/

	s, ok := i.(string)
	fmt.Println(s, ok)

	/*
		atribution doesnt break because of the type assertion.
		i := i.(float64) would break.
	*/
	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*
		It is possible to create a type switch, to check several assertions
	*/

	switch q := i.(type) {
	case int:
		fmt.Printf("%v is an int\n", q)
	case string:
		fmt.Printf("%v is a string\n", q)
	default:
		fmt.Printf("I don't know the type of %v\n", q)
	}

}
