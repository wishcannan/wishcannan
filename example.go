package main

import (
	"github.com/kataras/iris/v12"
	"services"
	//"encoding/json"
	"fmt"
)

func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
	    ctx.JSON(iris.Map{"message": "pong"})
	})

	app.RegisterView(iris.HTML("./home", ".html"))

	app.HandleDir("/public","./home/public")

	app.Get("/", func(ctx iris.Context) {
		ctx.View("樱花庄の宿舍.html")
	})

	app.Get("/home.html", func(ctx iris.Context) {
		ctx.View("home.html")
	})

	app.Get("/aboutme.html", func(ctx iris.Context) {
		ctx.View("aboutme.html")
	})
	app.Get("message.html", func(ctx iris.Context) {
		ctx.View("message.html")
	})
	app.Get("wushiyin.html",func(ctx iris.Context){
		ctx.View("wushiyin.html")
	})
	app.Get("/weather", func(ctx iris.Context) {
		w := services.GetWeather("三亚")
		    //fmt.Println(w)
		ctx.ViewData("temperature", w.Temperaturer)
		ctx.View("test.html")
		    //ctx.Writef("%s天气，%d摄氏度, %s, %s, %s" , w.Place, w.Temperaturer, w.Other1, w.Other2, w.Update_time)
	})
	app.Get("/api/weather", func(ctx iris.Context) {
		w := services.GetWeather("三亚")
		ctx.JSON(w)
	})
	app.Get("/api/view/{title:string}", func(ctx iris.Context){
		title := ctx.Params().Get("title")
		view := services.GetView(title)
		ctx.Writef("%d",view)
	})

	app.Get("api/getallmessage",func(ctx iris.Context){
		msgs := services.GetAllMsg()
		ctx.JSON(msgs)
	})

	app.Post("/api/liuyan", func(ctx iris.Context) {
		msg := ctx.FormValue("msg")
		fmt.Println(msg)
		services.WriteMsg(msg)
	})


					    // Listens and serves incoming http requests
					        // on http://localhost:8080.
	//app.Listen("127.0.0.1:8080")
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
