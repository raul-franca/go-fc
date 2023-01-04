package main

import "fmt"

// Com a seta a direita o canal só recebe valores
func receve(name string, ch chan<- string) {
	ch <- name
}

// Com a seta para esquerda o canal só passa o valor
func read(ch <-chan string) {
	fmt.Printf("Ola, %s", <-ch)
}
func main() {
	ch := make(chan string)
	go receve("Aurora", ch)
	read(ch)
}
