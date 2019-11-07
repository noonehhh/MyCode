package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string,2)
	ch <- "hello"
	ch <- "world"
	go testSeandFor(ch)
	time.Sleep(time.Second*2)
}

func testSeandFor(ch chan string){
	for{
		select {
		case v,ok := <- ch:
			if !ok{
				fmt.Println("close",v)

			}
			fmt.Println("value=:",v)
		}
	}
}