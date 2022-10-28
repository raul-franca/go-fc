package main

import (
	"fmt"
)

func main() {
	// closure = função anonima
	total := func() int {
		return soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
