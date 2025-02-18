package main

import "fmt"

func main() {
	func() { // anonymous function without parameters and return
		fmt.Println("Olá, mundo!")
	}() //"()" calls the function immediately
}
