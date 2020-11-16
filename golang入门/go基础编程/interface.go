package main

import "fmt"

type person struct {
	name string
	age  int
}

type people interface {
	eat(string)
}

func (p person) eat(s string) {
	fmt.Printf("%s is eating %s",p.name,s)
}

func main() {
	var p1 people
	p1=person{
		name: "jinlong",
		age:  18,
	}
	p1.eat("apple")
}