package main

import "fmt"

func main() {

	valor, err := sum(15, 10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, fmt.Errorf("a SOMA é maior que 49")
	}
	return a + b, nil
}
