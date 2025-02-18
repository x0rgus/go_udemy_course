package main

import "fmt"

func main() {
	func() { // anonymous function without parameters and return
		fmt.Println("Olá, mundo!")
	}() //"()" calls the function immediately

	func(texto string) { // anonymous function with parameter
		fmt.Println(texto)
	}("Olá, mundo! isso é um parametro!")

	retorno := func(texto string) string { // anonymous function with parameter and return
		return fmt.Sprintf("Recebido -> %s", texto) // formats the string and returns it
	}("Olá, mundo! isso é um parametro retornado!")
	fmt.Println(retorno)
}
