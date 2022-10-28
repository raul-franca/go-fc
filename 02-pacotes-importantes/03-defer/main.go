package main

func main() {
	defer println("Terceira linha")
	println("Primeira linha")
	println("Segunda linha")
}
