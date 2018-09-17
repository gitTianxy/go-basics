package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
	"os"
)

var rcv context.Handler
var reqLogger context.Handler
var logFile *os.File
var validate *validator.Validate

func init() {
	rcv = recover.New()
	reqLogger = logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	logFile, _ = os.OpenFile("web/iris.log", os.O_CREATE|os.O_WRONLY, 0666)
	validate = validator.New()
}

func main() {
	app := iris.Default()
	//app := iris.New()
	//app.Use(rcv)
	//app.Use(reqLogger)
	app.Logger().SetOutput(io.MultiWriter(logFile, os.Stdout))
	defer logFile.Close()

	simpleRequest(app)

	getDemo(app)

	userCRUD(app)

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}

func simpleRequest(app *iris.Application) {
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
}

func getDemo(app *iris.Application) {
	app.Get("/users/{id:long}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt64("id")
		ctx.Writef("get id: %v", id)
	})

	// len(name) <=10 otherwise this route will fire 404 Not Found
	// and this handler will not be executed at all.
	app.Get("/profile/{name:string max(10)}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("get name: %v", name)
	})

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe.
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		// shortcut for ctx.Request().URL.Query().Get("lastname").
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})
}

type User struct {
	Name string `json:"name" validate:"required,lt=255"`
	Age       uint8  `json:"age" validate:"gte=0,lte=130"`
	Sex     string `json:"sex" validate:"oneof=male female"`
	Email     string `json:"email" validate:"required,email"`
	Hobbies []string `json:"hobbies", validate:"required,dive,required"`
}

func userCRUD(app *iris.Application) {
	v1 := app.Party("/user")
	{
		v1.Get("/{id:long}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			ctx.Application().Logger().Infof("Get user. id:%v", id)
		})
		//validate.RegisterStructValidation(MyUserStructLevelValidation, User{})
		v1.Post("", func(ctx iris.Context) {
			var u User
			ctx.ReadJSON(&u)
			err := validate.Struct(u)
			if err != nil{
				ctx.StatusCode(http.StatusBadRequest)
				ctx.WriteString(err.Error())
			}
			ctx.Application().Logger().Info("POST user. body:", u)
		})
		v1.Put("/{id:long}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			var u User
			ctx.ReadJSON(&u)
			ctx.Application().Logger().Infof("PUT user. id:%v, new user:%v", id, u)
		})
		v1.Delete("/{id:long}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			ctx.Application().Logger().Infof("DELETE user. id:%v", id)
		})
	}
}

func myUserStructLevelValidation(sl validator.StructLevel)  {
	_ = sl.Current().Interface().(User)
}