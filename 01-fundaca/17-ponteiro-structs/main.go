package main

import "fmt"

type Conta struct {
	Saldo int
}

func (c Conta) simular(valor int) int {
	c.Saldo += valor
	return c.Saldo
}

func (c *Conta) addFundos(valor int) {
	c.Saldo += valor
}

func (c Conta) getSaldo() int {
	return c.Saldo
}

func main() {

	conta := Conta{Saldo: 100}
	fmt.Printf("Valor da simulação: R$ %d \n", conta.simular(400))
	fmt.Printf("Saldo da conta: R$ %d \n", conta.getSaldo())
	fmt.Println("Depositar R$ 400")

	conta.addFundos(400)
	fmt.Printf("Saldo da conta: R$ %d \n", conta.getSaldo())
}
