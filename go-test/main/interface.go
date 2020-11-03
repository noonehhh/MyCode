package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (p person) SetName(name string) {
	p.name = name
	fmt.Println("my name is ", p.name)
}

func (p *person) des() {
	fmt.Printf("my name is %v,i am %v", p.name, p.age)
}
func main() {
	p := person{"小明", 22, "man"}
	p.SetName("小花")
	p.des()
}
