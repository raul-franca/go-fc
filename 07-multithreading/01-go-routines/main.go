package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is runing \n", i, name)
		time.Sleep(time.Second / 2)
	}
}

// Thread 1
func main() {
	//Thread 2
	go task("A")
	//Thread 3
	go task("B")
	//Thread 3
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: task %s is runing \n", i, "C")
			time.Sleep(time.Second / 2)
		}
	}()

	//segura o processo -> Thread 1
	time.Sleep(time.Second * 10)
}
