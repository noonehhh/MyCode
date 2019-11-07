package main

import (
	"fmt"
)

func main(){

f("go")

go f("go1")

go func(msg string){
	fmt.Println(msg)
}("going")
//time.Sleep(1)
//var input string
//fmt.Scanln(&input)
//fmt.Println("done")

}
func f(from string){
	for i := 0;i < 3;i++{
      fmt.Println(from,":",i)
	}
}