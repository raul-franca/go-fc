package main

type MyNumber int

type number interface {
	~int | ~float64 // O ~ permite que o Go considere MyNumber como int
}

func soma[T int | float64](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func soma2[T number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"raul": 10, "João": 20, "Maria": 30}
	m2 := map[string]float64{"raul": 10.3, "João": 20.4, "Maria": 30.2}
	m3 := map[string]MyNumber{"raul": 10, "João": 20, "Maria": 30}

	println(soma(m))
	println(soma(m2))
	println("Compara: ", compara(10, 10))
	println(soma2(m))
	println(soma2(m2))
	println(soma2(m3))
}
