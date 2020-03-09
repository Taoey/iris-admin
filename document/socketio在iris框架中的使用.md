# socketio在iris框架中的使用

## 前言

本文中使用的依赖库是[go-socket.io](https://github.com/googollee/go-socket.io)，参考的是官网给出的例子：[iris示例](https://github.com/googollee/go-socket.io/blob/master/example/iris/main.go)

本文使用的socketio版本为 [v1.4.2](https://github.com/googollee/go-socket.io/tree/v1.4.2)

需要实现的功能：

- go-socketio-client：使用go语言实现对go-sockio的监听和事件的触发响应
- vue-socketio-client：使用vue实现socketio-client，实现事件的监听及触发响应





## 基于Iris创建一个socketio服务端

### socketio服务器构建

命名空间统一设置为`/`

```go
func SetWebsocketioServer()  {
	server, err := socketio.NewServer(nil)
	SocketServer = server
	if err != nil {
		log.Fatal(err)
	}
	SocketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	SocketServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	SocketServer.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
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
```

### 路由设置

其含义为：如果请求中包含`socket.io`请求接入到socketio服务器中，其他请求正常处理。

```go
// 设置socketio
App.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
    path := r.URL.Path
    fmt.Println(path)
    if strings.HasPrefix(path, "/socket.io/") {
        SocketServer.ServeHTTP(w, r)
        return
    }else {
        router(w, r)
        return
    }
})
```





## 基于golang-socketio的socketio客户端
socketio 客户端使用的是：[golang-socketio](https://github.com/graarh/golang-socketio)这个库

执行如下代码即可，不过比较可惜的是的没有找到命名空间的设置。

```go
	c, _ := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 9005, false), //ws://localhost:9005/socket.io/?EIO=3&transport=websocket
		transport.GetDefaultWebsocketTransport(),
	)

	c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})


	result, _ := c.Ack("msg","你好", time.Second*5)
	fmt.Println(result)

	c.Close()
```





基于Vue的socketio客户端使用



安装：

```
npm install vue-socket.io --save

npm install socket.io-client --save
```



使用：

在app.vue 中添加

```js
import VueSocketIO from '../../node_modules/vue-socket.io';

Vue.use(new VueSocketIO({
  debug: true,
  // connection: `http://${configs.server.host}:8080`,
  connection: 'http://127.0.0.1:8080',
  options: {
    path: '/socket.io/',
    forceNew: true,
    transports: ['websocket'],
    upgrade: false
  } //Optional options
}));
```



进行事件的监听，函数的名字就是要监听的事件名

```js
  sockets: {
    connect() {
      console.log('连接成功', this.$socket.id);
      this.$socket.emit('httpclient');
    },
    boxDisconnect(data) {
      data = JSON.parse(data);
      ElNotification({
        type: 'error',
        title: '离线报警',
        message: `设备已离线：${data.boxNo}-${data.boxName}`,
        duration: 0
      });
    },
    disconnect() {
      console.log('连接断开', this.$socket.id);
    }
  },
```



