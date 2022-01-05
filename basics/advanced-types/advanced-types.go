package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		Pointers
		A pointer holds the memory adress of a value
	*/

	value := 10
	pointer := &value // point to address
	*pointer = 12     // changing the value through the pointer

	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Pointer: %v\n", pointer)
	fmt.Printf("Pointer value: %v\n", *pointer)

	//--------------------------------------------------------------------------------------------------------------

	/*
		Structs
		A struct is a collection of fields, similar to classes.
		Structs that start with uppercase are available to other packages.
		Like classes, the struct field can be accessed using a dot
	*/

	type Person struct {
		FirstName string
		LastName  string
		secret    string
	}

	user := Person{"Daniel", "Freire", "Likes cats"}

	fmt.Println(user)
	fmt.Printf("Secret : %v\n", user.secret)

	structPointer := &user

	/*
		Following the pointers logic, to access a value of a struct pointer you
		would user (*pointer).field , but Go simplifies it to p.field
	*/

	fmt.Printf("Secret through pointer: %v\n", structPointer.secret)

	//--------------------------------------------------------------------------------------------------------------

	/*
		Arrays
		Arrays in Go have a fixed lenght that should be declared on instantiation.
		The syntax to create an array is : [<length>]T{<values>},
		where T is a type, length is an integer and values are the initial array values in order.
		Values that are not initialized are defined as the "zero value" for the type
	*/

	stringArray := [2]string{"element1", "element2"}

	fmt.Printf("String array : %v\n", stringArray)

	var intArray [3]int
	intArray[0] = 0
	intArray[1] = 1

	fmt.Printf("Int array : %v\n", intArray)

	//--------------------------------------------------------------------------------------------------------------

	/*
		Slices
		Slices are a dinamically sized view into the elements of an array.
		A slice is specified by: a[low : high] where a is an array reference,
		low is a lower bound index and high is a high bound index.
		The high index is excluded.

	*/

	baseArray := [6]int{2, 3, 5, 7, 11, 13}

	arraySlice := baseArray[1:4]
	fmt.Printf("Original array : %v\n", baseArray)
	fmt.Printf("Array slice : %v\n", arraySlice)

	/*
		Slices are like references do arrays.
		Changing the elements of a slice modify the elements of the original array
	*/

	arraySlice[0] = 1

	fmt.Printf("Original array after slice modification: %v\n", baseArray)

	/*
		You can create a slice with same syntax of an array but ommiting the lenght
	*/
	slice := []int{1, 2, 3, 4}

	fmt.Printf("Slice : %v\n", slice)

	/*
		Slices have lenght and capacity.
		The lenght represents the number of elements it contains.
		The capacity is the number of elements in the underlying array counting
		from the first element of the slice.
		The lenght and capacity of a slice can be obtained using the len() and cap() expressions.
	*/

	slice = slice[0:2]

	fmt.Printf("Slice Len : %v, Slice Cap: %v\n", len(slice), cap(slice))

	/*
		The zero value of a slice is nil.
		A slice is nil when the lenght is 0 and has no underlying array
	*/

	var nilslice []int
	fmt.Printf("Nilslice len : %v, Nilslice cap : %v\n", len(nilslice), cap(nilslice))
	if nilslice == nil {
		fmt.Println("Slice is nil")
	}

	/*
		Slices can be created with the make() function.
		Slices created with make are based on zeroed arrays.
	*/

	makeSlice := make([]int, 5)
	fmt.Printf("Make slice : %v\n", makeSlice)

	/*
		You can create slices of slices, they work like a matrix
	*/

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println("Slice board:")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/*
		Since slices have dinamic length you can add new elements to it.
		To add new elements to a slice use the append() function.
	*/

	newSlice := []string{}
	fmt.Printf("New slice : %v \n", newSlice)
	newSlice = append(newSlice, "element1", "element2")
	fmt.Printf("Slice after append : %v \n", newSlice)

	/*
		When using for loops with slice it is possible to range over the slice.
		Two values are returned for every iteration, the index and a copy of the element at that index
	*/

	for j, r := range newSlice {
		fmt.Printf("id : %d , value : %s \n", j, r)
	}

	/*
		To skip the index or value of a range you can use _.
	*/

	for _, r := range newSlice {
		fmt.Printf("value : %s \n", r)
	}

	//-------------------------------------------------------------------------------------------------------------

	/*
		Maps
		Maps are key/value pairs just like in other languages.
		Maps are declared with the syntax : map[key-type]value-type.
		Maps can also be created with the make function.
		To add a key/value pair to a map we use a syntax similar to arrays:
		m[key] = value, where m is a map variable.
	*/

	var m map[int]string
	m = make(map[int]string)
	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"

	fmt.Printf("Map : %v\n", m)

	/*
		To delete an element from a map you can use the delete function
		delete(map, key)
	*/

	delete(m, 0)

	fmt.Printf("Map after delete: %v\n", m)

	/*
		To test if a key is present use the syntax :
		elem, ok = m[key]
		If key is in m, ok is true. If not ok is false.
		If key is not in the map, then elem is the zero value for the map's element type
	*/

	elem, ok := m[1]
	fmt.Printf("Is 1 present ? %v, with value %v\n", ok, elem)
}
