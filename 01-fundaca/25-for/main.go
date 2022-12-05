package main

func main() {

	println("ex: 01 basico")
	for i := 0; i <= 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5} // slice

	println("ex: 02-CRUD range com indice")
	for i, numero := range numeros {
		println(i, numero)
	}

	println("ex: 03 range sem indice")
	for _, numero := range numeros { // "_" oculta o indice
		println(numero)
	}

	println("ex: 04 como um while")
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	println("ex: 05-percorrendo-arrays loop infinito")
	println("for { ... }")
	// funciona e é permitido no Go
	//for {
	//	println("Hello.world!")
	//
	//}
}
