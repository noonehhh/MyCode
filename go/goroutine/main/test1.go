package main

import (
	"fmt"
	"sync"
	"time"
)

/**
使用sync.WaitGroup阻塞主线程的执行，直到所有的goroutine执行完成
 */
func main(){
	var wg sync.WaitGroup
	for i := 0;i < 5;i++{
		wg.Add(1)
		go func(n int){
			defer wg.Add(-1)
			sss(n)
		}(i)
	}
	wg.Wait()
}
func sss(i int){
time.Sleep(3e9)
fmt.Println(i)
}