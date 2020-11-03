package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	app.Get("/getreq", func(context iris.Context) {
		path := context.Path()
		//app.Logger()
		fmt.Println(path)
	})

}
