package main

import "fmt"

// soma is a variadic function that accepts a variable number of integers and returns their sum.
func soma(numeros ...int) int {
	soma := 0
	// Iterate over each number in the variadic parameter and add it to the sum.
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// escrever is a function that accepts a string and a variable number of integers.
// It prints the string followed by each integer.
func escrever(texto string, numeros ...int) {
	// Iterate over each number in the variadic parameter and print it with the provided text.
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Call the soma function with different sets of integers and print the results.
	fmt.Println(soma(1, 2, 3))       // Output: 6
	fmt.Println(soma(1, 2, 3, 4, 5)) // Output: 15

	// Call the escrever function with a string and a set of integers.
	escrever("Olá", 1, 2, 3, 4, 5)
}
