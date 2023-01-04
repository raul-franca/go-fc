package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(11)

	go publish(ch)
	go read(ch, &wg)

	wg.Wait()
}

func read(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	//close(ch)
}
