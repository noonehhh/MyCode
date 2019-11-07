package main

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
)

func main(){
    a,b,c, _ := host.PlatformInformation()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
