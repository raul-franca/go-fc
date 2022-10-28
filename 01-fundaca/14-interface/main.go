package main

import (
	"fmt"
)

type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Empresa struct {
	Nome  string
	Ativo bool
}

func (e *Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %s foi desativada\n", e.Nome)

}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado %t \n", c.Nome, c.Ativo)

}

func Desativacao(p Pessoa) {
	p.Desativar()

}

func main() {
	raul := Cliente{
		Nome:  "Raul",
		Idade: 34,
		Ativo: true,
	}
	raul.Cidade = "Recife"
	fmt.Printf("O cliente %s tem %d mora em %s e esta %t \n", raul.Nome, raul.Idade, raul.Cidade, raul.Ativo)

	Desativacao(&raul)
	fmt.Printf("O cliente %s esta %t \n", raul.Nome, raul.Ativo)

}
