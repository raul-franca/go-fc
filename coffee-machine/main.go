package main

import (
	"fmt"
)

func main() {

	var coffeeCops = 0
	var water int
	var milk int
	var beans int

	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&coffeeCops)
	fmt.Printf("For %d cups of coffee you will need:\n", coffeeCops)
	water = 200 * coffeeCops
	milk = 50 * coffeeCops
	beans = 15 * coffeeCops
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)

	//5000 ml of water
	//1250 ml of milk
	//375 g of coffee beans

}
