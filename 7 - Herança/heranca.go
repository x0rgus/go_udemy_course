package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}
type estudante struct {
	pessoa    // passes all fields from 'pessoa' struct to 'estudante' struct (composition)
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo heranca") // Printing a message to indicate the file
}
