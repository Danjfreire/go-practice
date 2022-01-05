package main

import (
	"fmt"
)

/*
	You can declare multiple variables of the same type at once.
*/
var x, y, z bool

/*
	You don't need to declare a type of a variable when it is declared and initialized
*/
var a, b, c = 1, 2, 3

func main() {
	fmt.Println(x, y, z)
	fmt.Println(a, b, c)

	/*
		Inside functions you can use the shorthand operator := to declare and initialize
		a variable, ommiting the var keyword
	*/
	k := 10

	fmt.Println(k)
}
