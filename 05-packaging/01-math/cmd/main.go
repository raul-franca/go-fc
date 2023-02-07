package main

import (
	"fmt"
	math "github.com/raul-franca/go-fc/05-packaging/01-math"
)

func main() {
	var m math.Math
	m.A = 2
	m.B = 3
	fmt.Println(`m.A = 2 m.B = 3`)
	fmt.Printf("O resultado de m.add() = %v", m.Add())
}
