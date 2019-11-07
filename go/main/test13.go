package main

import (
	"fmt"
	"os"
	"reflect"
)

func main(){
	//date := time.Now()
	//day,_ := time.ParseDuration("-24h")
	//date2 := date.Add(day)
	//str := date.String()[:10]
	//fmt.Println(date)
	//fmt.Println(str)
	//fmt.Println(date2)
	//str := "./log/smart.log." + year + month + day
	//fmt.Println(time.Now().String()[:10])
	f, _ := os.Open("./log/smart.log.")
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

}
