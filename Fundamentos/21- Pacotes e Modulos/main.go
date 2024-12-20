package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)

	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Marca)
	fmt.Println(s)
	fmt.Println(matematica.A)
}
