package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops in Go")
	// Go only has one looping construct, the for loop
	// The for loop has three components: the init statement, the condition, and the post statement
	// The init statement will execute before the first iteration of the loop
	// The condition is evaluated before every iteration of the loop
	// The post statement is executed at the end of every iteration of the loop
	// The init statement, condition, and post statement are all optional
	// The for loop can be used to iterate over arrays, slices, maps, and strings
	// The for loop can also be used to create infinite loops

	// Basic for loop
	i := 0
	for i < 10 {
		time.Sleep(time.Second / 10)
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ { // The init statement, condition, and post statement are all optional
		time.Sleep(time.Second / 10)
		fmt.Println(j)
	}

	nomes := [3]string{"João", "Maria", "José"}

	for index, nome := range nomes { // The for loop can be used to iterate over arrays, slices, maps, and strings
		fmt.Println(index, nome)
	}

	for _, nome := range nomes { //The index can be ignored using the blank identifier (_)
		fmt.Println(nome)
	}

	for index, letra := range "ABC" { // The for loop can be used to iterate over arrays, slices, maps, and strings
		fmt.Println(index, string(letra)) // type conversion is required to convert the integer to a string, otherwise the ASCII value will be printed
	}

	usuarios := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	for key, value := range usuarios { // The for loop can be used to iterate over arrays, slices, maps, and strings
		fmt.Println(key, value)
	}
	for {
		// Infinite loop
		// The for loop can also be used to create infinite loops
		time.Sleep(time.Second / 10)
		fmt.Println("Infinite loop")
	}

}
