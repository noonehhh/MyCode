package main

import "fmt"

func main(){
	n := 0
	add(&n)
	fmt.Println(n)
}
func add(n *int){
	*n= *n+1
}