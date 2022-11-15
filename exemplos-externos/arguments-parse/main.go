package main

import (
	"fmt"
	"os"
)

// cmd: go run main.go Raul França Aurora

func main() {
	args := os.Args
	fmt.Println(args)
	println(len(args))

	for i, arg := range args {
		// Desconsiderar o primeiro argumento
		if i != 0 {
			fmt.Printf("valor: %v, %T \n", arg, arg)
		}
	}

}
