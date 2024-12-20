package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T Number](a T, b T) bool {
	if a > b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Thauan": 1500, "Rodrigues": 1515, "Mariazinha": 9090}

	m2 := map[string]float64{"Thauan": 15.30, "Rodrigues": 15.10, "Mariazinha": 90.90}

	m3 := map[string]MyNumber{"Thauan": 1500, "Rodrigues": 1515, "Mariazinha": 9090}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.0))
}
