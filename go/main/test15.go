package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	flist, e := ioutil.ReadDir("./main")
    fmt.Println(len(flist))
	if e != nil {
		fmt.Println("Read file error")
		return
	}

	//for _, f := range flist {
	//	if f.IsDir() {
	//		fmt.Println(tab, "+", dirPath+"/"+f.Name())
	//		readDir(dirPath+"/"+f.Name(), tab+"\t") //一股浓浓的函数编程。
	//	} else {
	//		fmt.Println(tab, ".", dirPath+"/"+f.Name())
	//	}
	//
	//}
}
