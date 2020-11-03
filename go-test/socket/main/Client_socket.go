package main

import (
	"fmt"
	"net"
	"time"
)

const RECV_BUF_LENC = 1024

func checkc(err error) {
	if err != nil {
		panic("this err is:" + err.Error())
	}
}
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:6666")
	checkc(err)
	defer conn.Close()
	buf := make([]byte, RECV_BUF_LENC)
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		fmt.Println(msg)
		//从服务器端收字符串
		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		fmt.Println(string(buf[0:n]))

		//等一秒钟
		time.Sleep(time.Second)
	}
}
