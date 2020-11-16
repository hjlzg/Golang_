package main

import "fmt"

func main() {
	var m=make(map[int]int)
	m[0] = 1
	m[2]=2
	for k,v:=range m{
		fmt.Printf("k=%v,v=%v",k,v)
		fmt.Println()
	}
}