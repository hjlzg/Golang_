package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	go func() {
		fmt.Printf("GN:%d\n", runtime.NumGoroutine())
		runtime.Goexit()
		fmt.Println("Exited")
	}()
	runtime.Gosched()
	fmt.Println("Ended.")
	fmt.Printf("1")
}
