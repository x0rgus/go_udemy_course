package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}
type estudante struct {
	pessoa    // passes all fields from 'pessoa' struct to 'estudante' struct (composition)
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo heranca") // Printing a message to indicate the file
	p1 := pessoa{"joão", "pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "USP"} // Creating a 'estudante' struct instance with a 'pessoa' instance
	fmt.Println(e1)
	fmt.Println(e1.nome) // Accessing the 'nome' field from the 'pessoa' struct

}
