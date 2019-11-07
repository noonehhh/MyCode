package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func cal(a int , b int )  {
	//defer wg.Add(-1)
	defer wg.Done()
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}

func main() {

	for i :=0 ; i<10 ;i++{
		wg.Add(1)
		//defer wg.Done()
		go cal(i,i+1)  //启动10个goroutine 来计算
	}
	wg.Wait()
	//time.Sleep(time.Second * 2) // sleep作用是为了等待所有任务完成
}
