package main

import "fmt"

type Client struct {
	nome string
}

func NewClient(nome string) *Client {
	return &Client{nome: nome}
}
func (c *Client) andou() {
	c.nome = "Thauan Rodrigues"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	thauan := Client{
		nome: "thauan",
	}
	thauan.andou()
	fmt.Printf("O valor da struct com nome é %v\n", thauan.nome)
}
