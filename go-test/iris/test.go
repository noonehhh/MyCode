package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/getreq", func(context iris.Context) {
		path := context.Path()
		//控制台打印路径
		app.Logger().Info(path)
		//写入返回数据
		context.WriteString("hello")
	})

	app.Get("/userinfo", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//接受前端返回的参数
		username := context.URLParam("username")
		app.Logger().Info(username)
		pwd := context.URLParam("password")
		app.Logger().Info(pwd)
	})
	//post方法
	app.Post("/postinfo", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//post请求向获取参数
		username := context.PostValue("username")
		app.Logger().Info(username)

		pwd := context.PostValue("password")
		app.Logger().Info(pwd)
	})
	type Person struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	//json数据的处理
	app.Post("/jsoninfo", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)

		var person Person
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		fmt.Println(person.Name)
		fmt.Println(person.Password)
		//data,_ := json.Marshal(&person)
		//json.Unmarshal(data,&person)

		//context.JSON(&person)
		//app.Logger().Info(person.name,person.pssword)

	})
	//正则表达式处理动态变量
	app.Get("/info/{name}/{age}", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		name := ctx.Params().Get("name")
		age := ctx.Params().Get("age")
		fmt.Println(name, age)
	})
	//正则表达式限制参数返回类型
	app.Get("/infoapi/{islogin:bool}", func(ctx iris.Context) {
		islogin, err := ctx.Params().GetBool("islogin")
		if err != nil {
			ctx.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if islogin {
			fmt.Println("已登录")
		} else {
			fmt.Println("未登录")
		}
	})
	app.Run(iris.Addr(":8900"))
}
