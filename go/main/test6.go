package main

import (
	"fmt"
	"time"
)

func main(){
sli :=make([]interface{},5)
sli[0] = 10
sli[1] = "haha"
sli[2] = true
sli[3] = 3.14
sli[4] = time.Now()

for i,v := range sli{
	fmt.Println(i,v)
}

}

