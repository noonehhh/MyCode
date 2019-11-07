package main

import (
	"fmt"
	"io"
	"net"
)

const RECV_BUF_LEN = 1024

func check(err error ){
	if err != nil{
		panic("this err is:"+err.Error())
	}
}

func main(){
listener,err := net.Listen("tcp","0.0.0.0:6666")//监听6666端口
if err != nil {
	panic("err listener:"+err.Error())
}
fmt.Println("start the server")
for{
	conn ,err :=listener.Accept()//接受连接
   check(err)
fmt.Println("Accepted the Connection",conn.RemoteAddr())
go EchoServer(conn)
}
}
 func EchoServer(conn net.Conn){
 	buf := make([]byte,RECV_BUF_LEN)
 	defer conn.Close()
 	for{
 		n,err := conn.Read(buf)
 		//check(err)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Printf("warning:end of data:%s\n",err)
		default:
			fmt.Printf("Error:Reading data:%s\n",err)
			return
		
		}
	}
 }