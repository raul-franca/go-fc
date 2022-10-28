package main

import (
	"fmt"
)

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	raul := Client{Nome: "Raul", Idade: 34, Ativo: true}
	fmt.Printf("O cliente %s tem %d e esta %t", raul.Nome, raul.Idade, raul.Ativo)
}
