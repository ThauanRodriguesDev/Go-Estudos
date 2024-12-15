package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	*b = 60
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	fmt.Println(minhaVar1)
	fmt.Println(minhaVar2)
	fmt.Println(soma(&minhaVar1, &minhaVar2))
	fmt.Println(minhaVar1)
	fmt.Println(minhaVar2)
}
