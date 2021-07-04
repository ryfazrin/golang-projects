package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	beego.Get("/home",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Run()
}