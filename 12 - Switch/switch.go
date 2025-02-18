package main

import "fmt"

func diaDaSemana(numero int) string { // Function that returns the day of the week according to the number passed as a parameter
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough // The 'fallthrough' keyword is used to execute the next case
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana(1)
	fmt.Println(dia2) // The 'fallthrough' keyword makes the 'diaDaSemana' variable receive the value of the next case
}
