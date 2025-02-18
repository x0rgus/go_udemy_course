package main

import "fmt"

func soma(numeros ...int) int { // variadic function that accepts a variable number of integers
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	fmt.Println((soma(1, 2, 3))) // 6
}
