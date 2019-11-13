package main

import (
	"github.com/kataras/iris"
	"github_com_Taoey_iris_cli/src/modules/myapi"
	"github_com_Taoey_iris_cli/src/myinit"
	"log"
)

var App *iris.Application

//程序入口
func main() {
	// 初始化
	myinit.InitConf()
	//myinit.InitMongo()
	//myinit.InitQuartz()

	// 初始化App
	App = iris.New()
	SetRoutes()
	// 启动
	url := myinit.GCF.UString("server.url")
	run := App.Run(iris.Addr(url), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}

// 设置路由
func SetRoutes() {

	//主页
	App.Get("/", myapi.Index)
	App.Get("/hello_json", myapi.IndexHelloJson)

	//根API
	RootApi := App.Party("api/")

	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), myapi.UploadAliBill)

	RootApi.Get("/download/demo1", myapi.ExcelDownloadDemo1)
	RootApi.Get("/download/demo2", myapi.ExcelDownloadDemo2)
	RootApi.Get("/download/demo3", myapi.ExcelDownloadDemo3)
	RootApi.Get("/download/demo4", myapi.DownloadLimite)
	RootApi.Get("/download/demo5", myapi.SendURLFile) //http://localhost:9005/api/download/demo4

	RootApi.Get("/redirect/", myapi.RedirectURL)

	// 读取post信息
	RootApi.Post("/post/jsondemo", myapi.PostAcceptJson)
}
