package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"João", 25}
	fmt.Println(usuario2)

	// para declarar um struct sem passar todos os valores é necessario usar argumentos nomeados
	usuario3 := usuario{nome: "Maria"}
	fmt.Println(usuario3)
}
