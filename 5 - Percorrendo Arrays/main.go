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
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
