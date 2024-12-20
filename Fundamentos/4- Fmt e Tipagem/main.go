package main

import "fmt"

const a = "Hello-World"

type ID int

var (
	b bool    = true
	c int     = 12
	d float64 = 1.2
	e string  = "Thauan"
	f ID      = 1
)

func main() {
	//%T retorna o tipo utilizado e o %v retorna o valor
	fmt.Printf("O tipo padrão é %T", f)
}
