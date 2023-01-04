package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is runing \n", i, name)
		time.Sleep(time.Second / 2)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)
	//Thread 2
	go task("A", &waitGroup)
	//Thread 3
	go task("B", &waitGroup)
	//Thread 3
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: task %s is runing \n", i, "C")
			time.Sleep(time.Second / 2)
			waitGroup.Done()
		}
	}()
	//segura o processo -> Thread 1
	waitGroup.Wait()

}
