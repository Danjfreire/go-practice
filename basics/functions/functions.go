package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println((swap("world", "hello")))
	fmt.Println((split(8)))

	fmt.Println(compute(add))

	closure := closureFunc()

	fmt.Printf("Closure func after 1 execution : %v\n", closure())
	fmt.Printf("Closure func after 2 executions : %v\n", closure())
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

/*
	Functions can be passed around just like other values.
	Functions may be used as functions arguments and return values.
*/

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

/*
	Functions may be closures.
	A closure is a function value that references variables from outside its body.
	You can say that the function is "bound" to those variables
*/

func closureFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
