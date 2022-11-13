package main

import (
	"fmt"
	"time"
)

// Empregado Definição da struct
type Empregado struct {
	id       int
	nome     string
	pais     string
	inclucao time.Time
}

func main() {
	//declaração das variáveis
	var emp Empregado
	newEmp := new(Empregado) // new, cria um novo vazio

	// set valores
	emp.id = 1
	emp.nome = "Aurora"
	emp.pais = "Canada"
	emp.inclucao = time.Now()

	newEmp.id = 2
	newEmp.nome = "Aracelly"
	newEmp.pais = "Italia"
	newEmp.inclucao = time.Now()

	// display
	fmt.Println("=====================")
	fmt.Println(emp.id)
	fmt.Println(emp.nome)
	fmt.Println(emp.pais)
	fmt.Println(emp.inclucao)
	fmt.Println("=====================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.nome)
	fmt.Println(newEmp.pais)
	fmt.Println(newEmp.inclucao)

}
