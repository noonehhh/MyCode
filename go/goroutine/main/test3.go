package main

import "fmt"

func main(){
	n := 0
	ch := make(chan int,10)
	for i := 0;i < 10;i++{
		go testch(ch,&n)
	}
	for i := 0;i<10;i++{
		msg := <-ch
		fmt.Println(msg)
	}
	close(ch)
	fmt.Println("helloworld")

}

func testch(ch chan int,n *int){
	ch <- *n+1
}
