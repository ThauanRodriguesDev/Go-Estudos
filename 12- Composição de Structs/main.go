package main

import "fmt"

type Endereço struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereço
}

func main() {

	thauan := Cliente{
		Nome:  "Thauan",
		Idade: 30,
		Ativo: true,
	}
	thauan.Cidade = "Goiás"

	fmt.Println(thauan.Cidade)
}
