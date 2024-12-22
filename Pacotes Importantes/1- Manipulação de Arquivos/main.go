package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, error := f.Write([]byte("Escrevendo algo no arquivo"))

	if error != nil {
		panic(error)
	}

	fmt.Printf("Arquivo Criado com sucesso, tamanho: %d caracteres \n", tamanho)
	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
