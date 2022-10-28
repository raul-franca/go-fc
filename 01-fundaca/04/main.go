package main

import "fmt"

const h = "Hello, world!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Raul"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo da variavel c é %T", c)
}
