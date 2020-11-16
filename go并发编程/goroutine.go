package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("hello,I'm %s\n ", name)
}

func main() {
	name := "jinlong"
	go greet(name)
	// time.Sleep(10 * time.Millisecond)
	fmt.Printf("GoodBye %s\n", name)
}
