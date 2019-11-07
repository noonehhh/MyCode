package main

import (
	"fmt"
	"strings"
)

func main(){
//a := 3
//strconv.FormatInt(int64(a),10)
//fmt.Println(reflect.TypeOf(a))
//	tName := strconv.FormatInt(time.Now().UnixNano(), 10)+"LLL"+strconv.Itoa(rand.Intn(1000)) + ".xml"
//    fmt.Println(tName)
//    url := "local://"
//	realFile := "./files/" + strings.Replace(url, "local://", "REALDATA/", -1)
//	fmt.Println(realFile)
//ma :=map[string]string{"table1":"   hello"}
//fmt.Println(ma["table1"])
str :="hello"
s := strings.Replace(str,"he","..",-1)
fmt.Println(s)
}