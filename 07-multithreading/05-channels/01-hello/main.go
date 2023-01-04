package main

import "fmt"

// thread 1
func main() {
	canal := make(chan string) //vazio

	//thread 2
	go func() {
		canal <- "Ola mundo!" // cheio
	}()

	//thread 1
	msg := <-canal
	fmt.Print(msg)
}
