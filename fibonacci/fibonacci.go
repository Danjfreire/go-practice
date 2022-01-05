package main

import (
	"fmt"
)

var precalculated = make(map[int]int)

func main() {
	targetNum := 50
	fmt.Println("This is a fibonacci program")
	result := fibonacci(targetNum)
	fmt.Printf("Fib of %v is %v! \n", targetNum, result)
}

func fibonacci(x int) int {

	if len(precalculated) >= x && x > 1 {
		return precalculated[x]
	}

	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	res := fibonacci(x-1) + fibonacci(x-2)
	precalculated[x] = res

	return res
}
