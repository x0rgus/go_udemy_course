package main

import "fmt"

func CalculosMatematicos(a, b int) (soma, subtracao int) {
	soma = a + b
	subtracao = a - b
	return // implicit return of the named return values
}
func main() {
	soma, subtracao := CalculosMatematicos(10, 5)
	fmt.Println(soma, subtracao)

}
