package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// Pointers are variables that store memory addresses
	var variavel3 int
	var ponteiro *int
	variavel3 = 100

	ponteiro = &variavel3 // The '&' operator is used to get the memory address of a variable

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150                   // Changing the value of the variable
	fmt.Println(variavel3, *ponteiro) // The '*' operator is used to get the value of a memory address
}
