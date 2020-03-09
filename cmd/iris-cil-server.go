package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/kataras/iris"
	"github_com_Taoey_iris_cli/src/modules/myapi"
	"github_com_Taoey_iris_cli/src/myinit"
	"net/http"
	"strings"
)

var App *iris.Application
var SocketServer *socketio.Server

//程序入口
func main() {
	App = iris.New()
	// 初始化
	myinit.InitConf()
	//myinit.InitMongo()
	//myinit.InitQuartz()
	SetWebsocketioServer()
	// 初始化App
	SetRoutes()
	// 设置视图
	App.RegisterView(iris.HTML("./web/views", ".html"))
	// 设置socketio
	App.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		fmt.Println(path)
		if strings.HasPrefix(path, "/socket.io/") {
			// 允许websocket跨域
			origin := r.Header.Get("Origin")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			SocketServer.ServeHTTP(w, r)
			return
		} else {
			router(w, r)
			return
		}
	})
	// 启动
	url := myinit.GCF.UString("server.url")
	run := App.Run(
		iris.Addr(url),
		iris.WithCharset("UTF-8"),
		iris.WithoutPathCorrection,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	fmt.Println(run)
}

// 设置路由
func SetRoutes() {

	//主页
	App.Get("/", myapi.Index)
	App.Get("/index", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	App.Get("/socket_test", func(ctx iris.Context) {
		ctx.View("socketio.html")
	})
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

func SetWebsocketioServer() {
	server, err := socketio.NewServer(nil)

	SocketServer = server
	if err != nil {
		fmt.Println(err)
	}
	SocketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected-id:", s.ID())
		// 连接成功后触发客户端事件
		return nil
	})

	SocketServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	SocketServer.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		fmt.Println("msg:", msg)
		return "服务端接收到了:" + msg
	})
	SocketServer.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	SocketServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	SocketServer.OnError("/", func(e error) {
		fmt.Println(e)
	})
	go SocketServer.Serve()
}
