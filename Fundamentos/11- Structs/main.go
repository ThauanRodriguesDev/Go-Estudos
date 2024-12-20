package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	thauan := Cliente{
		Nome:  "Thauan",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(thauan.Ativo)
}
