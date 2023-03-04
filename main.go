package main

import (

	// "github.com/kataras/iris"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/yuta-hanyu/kakeibo-api/src/app/controller/users"
	"github.com/yuta-hanyu/kakeibo-api/src/app/service"
	"github.com/yuta-hanyu/kakeibo-api/src/app/setups"
)

func main() {
	app := iris.New()

	// ミドルウェアの使用
	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// ログ記録（これも備え付きミドルウェア）
	ac := accesslog.File("./access.log")
	defer ac.Close()

	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())

	dbMap := service.InitDb()
	defer dbMap.Db.Close()

	app.Get("/ping", pong)

	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())

	// CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodPatch, iris.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Front-Version", "X-Front-Version-For-Sales"},
		AllowCredentials: true,
	})

	app.UseRouter(crs)

	// "/users/"から始まるURLを受け取った際の処理をグループ化
	users := app.Party("/users", crs).AllowMethods(iris.MethodOptions)
	mvc.Configure(users, setups.ConfigureUsersControllers)

	// ポートの指定
	app.Listen(":8080")
}

func setup(app *mvc.Application) {
	// ログを取得してくれる機能のDIもここで行う
	app.Register(accesslog.GetFields)

	// URLが "/users" から始まるリクエストを受け取った際に，
	// 以下の Controllerを使用させるという指示
	app.Handle(new(users.UsersController))
}

func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}

// func main() {
// 	app := iris.New()

// 	// ミドルウェアの使用
// 	app.Use(iris.Compression)
// 	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

// 	app.Get("/ping", pong).Describe("healthcheck")

// 	// ログ記録（これも備え付きミドルウェア）
// 	// ac := accesslog.File("./access.log")
// 	// defer ac.Close()

// 	// CORS
// 	// crs := cors.New(cors.Options{
// 	// 	AllowedOrigins:   []string{"*"},
// 	// 	AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodPatch, iris.MethodOptions},
// 	// 	AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Front-Version", "X-Front-Version-For-Sales"},
// 	// 	AllowCredentials: true,
// 	// })

// 	// app.UseRouter(ac.Handler)
// 	// app.UseRouter(recover.New())

// 	// app.UseRouter(crs)

// 	// // app.UseRouter(ac.Handler)
// 	// // app.UseRouter(recover.New())

// 	// service.InitDb()

// 	// usersAPI := app.Party("/users")
// 	// mw := mvc.New(usersAPI)
// 	// // fmt.Print(user)
// 	// mvc.Configure(usersAPI, setups.ConfigureWeightsControllers)

// 	// weightsAPI := app.Party("/weights")
// 	// mw := mvc.New(weightsAPI)
// 	// mw.Handle(new(weights.WeightController))
// 	// user := app.Party("/users")
// 	// mw := mvc.New(user)
// 	// mw.Handle(new(user.UsersController))

// 	// "/users/"から始まるURLを受け取った際の処理をグループ化
// 	// users := app.Party("/users")
// 	// mvc.Configure(users, setups.ConfigureUsersControllers)

// 	// ポートの指定
// 	app.Listen(":8080")
// }

// func pong(ctx iris.Context) {
// 	// fmt.Println("=====")
// 	ctx.WriteString("pong")
// }

// // func setup(app *mvc.Application) {
// // 	// Register Dependencies.
// // 	app.Register(
// // 		environment.DEV,         // DEV, PROD
// // 		database.NewDB,          // sqlite, mysql
// // 		service.NewGreetService, // greeterWithLogging, greeter
// // 	)

// // 	// Register Controllers.
// // 	app.Handle(new(controller.GreetController))
// // }
