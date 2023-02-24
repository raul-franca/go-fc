package main

import "fmt"

func main() {
	var num int
	fatorial := 1

	fmt.Println("Digite um número:")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Não é possível calcular o fatorial de um número negativo.")
	} else if num == 0 {
		fmt.Println("O fatorial de 0 é 1.")
	} else {
		for i := 1; i <= num; i++ {
			fatorial *= i
		}
		fmt.Printf("O fatorial de %d é %d.\n", num, fatorial)
	}
}
