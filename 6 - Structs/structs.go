package main

import "fmt"

// Defining the 'usuario' struct with fields 'nome', 'idade', and 'endereco'
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco // It is possible to create a struct inside another struct
}

// Defining the 'endereco' struct with fields 'logradouro' and 'numero'
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs") // Printing a message to indicate the file

	var u usuario
	u.nome = "Davi" // Assigning values to the 'usuario' struct fields
	u.idade = 21
	fmt.Println(u) // Printing the 'usuario' struct

	enderecoExemplo := endereco{"Rua dos Bobos", 0}  // Creating an 'endereco' struct instance
	usuario2 := usuario{"João", 25, enderecoExemplo} // Creating a 'usuario' struct instance with an 'endereco' instance
	fmt.Println(usuario2)                            // Printing the 'usuario' struct

	// To declare a struct without passing all values, it is necessary to use named arguments
	usuario3 := usuario{nome: "Maria"}
	fmt.Println(usuario3) // Printing the 'usuario' struct
}
