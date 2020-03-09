package test

import (
	"fmt"
	"github.com/graarh/golang-socketio"
	_ "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	_ "github.com/graarh/golang-socketio/transport"
	"runtime"
	"testing"
	"time"
)

func sendJoin(c *gosocketio.Client) {
	//result, err := c.Ack("msg","你好", time.Second*5)
	result, err := c.Ack("msg", "你好", time.Second*5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ack result:", result)
	}
}

func Test01(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 9005, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		fmt.Println(err)
	}

	c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		fmt.Println("Connected")
		c.Emit("/notice", "聊天")
	})

	c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		fmt.Println("Disconnected")
	})
	c.On(gosocketio.OnError, func(h *gosocketio.Channel) {
		fmt.Println("error")
	})

	time.Sleep(time.Second * 2)
	go sendJoin(c)

	select {}
	c.Close()

}
