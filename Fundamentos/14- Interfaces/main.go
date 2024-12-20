package main

import "fmt"

type Endereço struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar() // O Go so permite passar metodos em interfaces
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereço
}

type Empresa struct {
	Nome string
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente foi desativado\n")
}

func (e Empresa) Desativar() {
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	// thauan := Cliente{
	// 	Nome:  "Thauan",
	// 	Idade: 30,
	// 	Ativo: true,
	// }
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
}
