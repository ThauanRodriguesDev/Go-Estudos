package main

func main() {
	a := 1
	b := 1

	if a > b && a != 0 {
		println(a)
	} else {
		println(b)
	}

	switch b {
	case 1:
		b = 3
	case 2:
		b = 4
	case 3:
		b = 5
	}
	println(b)
}
