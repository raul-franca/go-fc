package main

import "fmt"

func fatorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {
	fmt.Printf("O fatorial é: %d", fatorial(3))
}
