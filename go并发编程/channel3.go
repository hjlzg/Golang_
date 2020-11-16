package main

import "fmt"

func send(ch chan<- int, number int) {
	ch <- number
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 2)
	go send(ch2, 1)

	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("CH1:%d\n", v1)
		case v2 := <-ch2:
			fmt.Printf("CH2:%d\n", v2)
		}
	}

	// fmt.Println("End.")
}
