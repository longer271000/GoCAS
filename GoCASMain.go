package main

import (
	config2 "GoCAS/config"
	"GoCAS/controller"
	"GoCAS/datasource"
	"GoCAS/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

func main()  {

	//创建
	app := newApp()
	//应用设置
	configation(app)

	//路由设置
	mvcHandle(app)

	config := config2.InitConfig()
	addre := "localhost:" + config.Port

	app.Run(
		iris.Addr(addre),								//端口号8090
		iris.WithoutServerError(iris.ErrServerClosed),  //无服务错误提示
		iris.WithOptimizations,							//对JSON数据序列号更快的配置
		)


}
//APP构建
func newApp()  *iris.Application{

	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")

	//注册静态资源
	//app.HandleDir("/static","./static")
	//app.HandleDir("/image","./static/image")
	//app.HandleDir("/js","./static/js")
	////注册视图文件
	//app.RegisterView(iris.HTML("./static",".html"))
	//app.Get("/", func(context *context.Context) {
	//	context.View("index.html")
	//
	//})
	return app

}
/***
项目设置
**/
func configation(app *iris.Application)  {
	//编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//错误配置-未知错误
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"drrmsg": iris.StatusNotFound,
			"msg": "Not fount",
			"data": iris.Map{},

		})

	})

	//错误设置-未知错误
	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"drrmsg": iris.StatusNotFound,
			"msg": "Not fount",
			"data": iris.Map{},
		})
	})
}
//MVC 架构模式
func mvcHandle(app *iris.Application)  {
	seesManager := sessions.New(sessions.Config{
		Cookie: "sessioncookie",
		Expires: 24* time.Hour,

	})
	engine := datasource.NewMysqlEngine()

	//模块功能
	loginService := service.NewLoginInfoService(engine)

	userCASRouter := mvc.New(app.Party("/user"))
	userCASRouter.Register(
		loginService,
		seesManager.Start,
		)
	userCASRouter.Handle(new(controller.LoginInfoController))

	//公共信息
	app.Handle("GET","/about", func(ctx iris.Context) {
		ctx.HTML("<h1>这是统一用户授权中心</h1>")
	})
}

