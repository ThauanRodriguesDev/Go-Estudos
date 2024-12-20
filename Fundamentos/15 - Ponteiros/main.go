package main

func main() {
	//Memória -> Endereço -> Valor

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a

	println(*b)
}
