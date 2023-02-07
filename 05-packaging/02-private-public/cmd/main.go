package main

import (
	"fmt"
	math02 "github.com/go-fc/05-packaging/02-private-public"
)

func main() {
	m := math02.NewMath(2, 3)
	fmt.Printf("result m.Add(): a:2 e b:3 = %v\n", m.Add())
	m.C = 20
	fmt.Printf("m.C é do %T, e tem o valor de %v\n", m.C, m.C)
	fmt.Println(math02.ConstPulica)
}
