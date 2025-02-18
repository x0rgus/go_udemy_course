package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 { // The variable 'outroNumero' is only available inside the 'if' block, it is not available outside the block
		fmt.Println("Maior que zero")
	} else {
		fmt.Println("Menor que zero")
	}

}
