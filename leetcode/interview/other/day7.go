package main

import (
	"fmt"
	"math/rand"
)

/**
n个数 1-n，随机n次,将这n个数输出
*/
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := len(arr)
	for i := 0; i < index; i++ {
		inx := rand.Intn(len(arr))
		fmt.Println(arr[inx])
		arr = append(arr[:inx], arr[inx+1:]...)
	}
}
