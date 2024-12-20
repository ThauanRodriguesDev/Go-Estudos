package matematica

//Tudo que começa com letra maiuscula esta Exportado, tudo que começa com letra minuscula nao esta exportado

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}
