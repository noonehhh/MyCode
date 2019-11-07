package main

import (
	"io"
	"os"
)

func main(){
os.Mkdir("C:/Users/97556/Desktop/go",1)
os.Create("C:/Users/97556/Desktop/go/1.txt")
file,_ := os.OpenFile("C:/Users/97556/Desktop/go/1.txt",os.O_APPEND,777)
io.WriteString(file,"hello")

file.Close()
}
