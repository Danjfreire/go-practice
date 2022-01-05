package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println((swap("world", "hello")))
	fmt.Println((split(8)))
}

/*
*	The parameter types comes after the parameter variable separated by space.
*	If consecutive params share the same type, you can ommit the type from all but the latest,
*	the  example bellow would become : (x, y int)
 */
func add(x int, y int) int {
	return x + y
}

/*
	Functions can return multiple results
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
	Function return may be named, they are treated as variabled defined at the top of the functions and returned with a "naked" return.
	This should only be used in very short functions
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
