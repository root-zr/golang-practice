package main

import (
	"connToDatabase/conf"
	"connToDatabase/datasource"

	"connToDatabase/route"
	"flag"

	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	return app
}

func main() {
	datasource.InitDB()
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.Sysconfig.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}
