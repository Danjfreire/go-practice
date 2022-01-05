package main

import (
	"fmt"
)

func main() {

	/*
		Defer
		A defer statement defers the execution of a function until the
		surrounding function returns.
		Defers are pushed into a stack so the last in gets executed first
	*/
	defer fmt.Println("This is printed after main() executes - 2")
	defer fmt.Println("This is printed after main() executes - 1")

	sum := 1

	/*
		For
		The init and post statements are optional.
		The semicolons can be removed as well and the for "becomes a while".
		An infinite loop can be expressed as for {}
	*/

	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)

	/*
		Switch
		Go switches only run the selected case, no need to add breaks.
		Switches evaluate from top to bottom.


	*/

	value := "test"

	switch value {
	case "hue":
		fmt.Println("Hue case")
	case "test":
		fmt.Println("Test case")
	default:
		fmt.Println("Default case")
	}

	/*
		Ifs and Switches can have a short initialization statement:
	*/

	switch test := "xd"; test {
	case "xd":
		fmt.Println("xd")
	default:
		fmt.Println("default")
	}

	/*
		A switch without condition is an option to write long if-else chains
	*/

	test2 := 2

	switch {
	case test2 > 1:
		fmt.Println("Greater than 1")
	case test2 < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("0 or 1")
	}
}
