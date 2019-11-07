package main

import "fmt"

type usr interface {
	run()
}
type job struct {}
type inter interface {
	say()
	nosay()
}

type user struct {
	name string
}

func (u *user)say(){
	fmt.Println(u.name)
}
func (u *user)nosay(){
	fmt.Println(u.name)
}

func main(){
	var s inter = &user{"xiaoming"}
	s.nosay()
	s.say()
}