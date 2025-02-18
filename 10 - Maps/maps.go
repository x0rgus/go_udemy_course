package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Davi",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Pedro",
		},
		"curso": {
			"faculdade": "Engenharia",
			"campus":    "campus 1",
		},
	} // Creating a map with a map as a value
	fmt.Println(usuario2)
	delete(usuario2, "nome") // Deleting a key-value pair from the map
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	} // Adding a key-value pair to the map
	fmt.Println(usuario2)
}
