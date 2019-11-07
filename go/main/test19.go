package main

import "fmt"

func main(){
arr := ".0.1.2.3.5."
fmt.Println(arr[:len(arr)-2])
}
