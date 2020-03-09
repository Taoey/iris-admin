package test

import (
	"fmt"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"testing"
)

func TestSocketioClient2(t *testing.T) {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("127.0.0.1", 8080, false),
		transport.GetDefaultWebsocketTransport(),
	)
	fmt.Println(err)
	c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		h.Emit("httpclient", "你好")
	})

	c.On("boxDisconnect", func(h *gosocketio.Channel) {

		fmt.Println("终端断开连接")
	})
}
