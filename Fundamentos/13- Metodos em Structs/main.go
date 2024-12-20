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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente foi desativado\n")
}

func main() {

	thauan := Cliente{
		Nome:  "Thauan",
		Idade: 30,
		Ativo: true,
	}
	thauan.Cidade = "Goiás"

	thauan.Desativar()
	fmt.Println(thauan.Cidade)
	fmt.Println(thauan.Ativo)
}
