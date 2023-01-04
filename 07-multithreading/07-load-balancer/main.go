package main

import (
	"fmt"
	"time"
)

func worker(id int, data chan int) {
	for x := range data {
		fmt.Printf("Worker: %d, data: %d \n", id, x)
		time.Sleep(time.Second / 2)

	}
}

func main() {
	data := make(chan int)
	QTDworkers := 1000

	for i := 0; i < QTDworkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1500; i++ {
		data <- i
	}
}
