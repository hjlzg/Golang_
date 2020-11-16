package main

import "fmt"

var ch = make(chan int)
var content string

func set() {
	content = "come on"
	<-ch
}

func main() {
	go set()
	ch <- 0 //"ch<-0"语句会阻塞，直至"<-ch"执行完毕
	fmt.Println(content)
}
