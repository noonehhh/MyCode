package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main(){
	app :=iris.New()
	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
	app.Get("/getreq",func(context context.Context){
		path := context.Path()
		//app.Logger()
		fmt.Println(path)
	})

}