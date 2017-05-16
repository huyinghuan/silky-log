package main

import (
	"silky-log/controller"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
)

func main() {
	app := iris.New()
	sessionAdapt := sessions.New(sessions.Config{})

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		sessionAdapt)

	API := app.Party("/api", func(ctx *iris.Context) { ctx.Next() })
	logCtrl := &controller.LogCtrl{}
	API.Post("/log", logCtrl.Post)
	API.Get("/log", logCtrl.GetAll)
	app.Listen(":6400")
}
