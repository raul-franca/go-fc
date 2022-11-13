package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	x, y int
	r    float64
}

// (c Circulo) define ser um method de Circulo
func (c Circulo) display() {

	fmt.Printf("x=%d,y=%d,r=%.2f\n", c.x, c.y, c.r)
}
func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
func (c Circulo) circumference() float64 {
	return 2 * math.Pi * c.r
}

// Altera os valores apenas dentro do method
func (c Circulo) moveToSemPonteiro(newX, newY int) {
	c.x = newX
	c.y = newY
	fmt.Println("dentro do methodo sem ponteiro")
	c.display()
}

// Altera os valores no objeto
func (c *Circulo) moveToComPonteiro(newX int, newY int) {
	c.x = newX
	c.y = newY
	fmt.Println("dentro do methodo com ponteiro")
	c.display()
}

func main() {
	c := Circulo{10, 5, 2.8}
	c.display()

	fmt.Printf("Area do circulo: %2.f \n", c.area())
	fmt.Printf("Circuferencia: %2.f \n", c.circumference())
	c.moveToSemPonteiro(5, 2)
	fmt.Println("###Direto do main:")
	c.display()
	c.moveToComPonteiro(40, 30)
	fmt.Println("###Direto do main:")
	c.display()

}
