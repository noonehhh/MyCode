package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
)
type res struct {
	Page int
	Fruit []string
}
type res2 struct {
	Page  int        `json:"page"`
	Fruit []string   `json:"fruit"`
}
func main(){
	res1D := &res{
		Page:   1,
		Fruit: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
}
