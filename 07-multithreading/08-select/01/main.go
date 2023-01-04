package main

import "time"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		println("Received: ", msg1)

	case msg2 := <-c2:
		println("Received: ", msg2)

	case <-time.After(time.Second * 3):
		println("Timeout")
	}

}
