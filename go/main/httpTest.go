package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello,世界!!")
}
func login(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "请登陆!!!!")
}
