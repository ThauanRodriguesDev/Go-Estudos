package main

import "fmt"

func main() {

	salarios := map[string]int{"Thauan": 1000, "Lanzinho": 1015, "Seila": 3041}

	delete(salarios, "Lanzinho")

	salarios["Golanzinho"] = 1010

	// sal := make(map[string]int)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
