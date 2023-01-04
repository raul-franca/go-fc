package main

func main() {
	forever := make(chan bool)

	// ex1  deadlock!
	//forever <- true

	// ex2  deadlock!
	//<-forever

	//ex3 funciona porque forever chan foi alterado em outra thread
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

}
