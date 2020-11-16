package main

import "fmt"

func main() {
	var num = 72
	printNum(num)
}

func printNum(num int) {
	switch num2 := num; {
	case num2 < 0:
		fmt.Println(num2)
		fallthrough
	case num2 > 60:
		fmt.Println(num2)
		fallthrough
	default:
		fmt.Println("finish")
	}
}
