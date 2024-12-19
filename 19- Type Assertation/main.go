package main

func main() {
	var minhaVar interface{} = "Thauan Rodrigues"

	println(minhaVar.(string))
	resultado, ok := minhaVar.(int)
	println(resultado)
	println(ok)

	resultado2 := minhaVar.(int)
	println(resultado2)
}
