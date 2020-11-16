package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Printf("Only once\n")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(num int) {
			once.Do(onceBody) //once.Do中的方法不能有参数
			// onceBody(num)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
