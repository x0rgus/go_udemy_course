package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")
	// Arrays are fixed-size data structures that store elements of the same type
	// two different ways to declare an array
	var array1 [5]string

	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // The size of the array is defined by the number of elements
	fmt.Println(array3)

	// Slices are dynamic-size data structures that store elements of the same type

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3)) // The type of the array is different from the type of the slice

	slice = append(slice, 18) // Adding an element to the slice
	fmt.Println(slice)

	slice2 := array2[1:3] // Creating a slice from an array
	fmt.Println(slice2)

	array2[1] = "posição alterada" // Changing the value of an array element
	fmt.Println(slice2)            // The slice is a reference to the array, so the change in the array is reflected in the slice
}
