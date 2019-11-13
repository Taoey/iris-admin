package myapi

import (
	"encoding/json"
	"github.com/kataras/iris"
	"log"
)

// 用于测试接收Json参数

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PostAcceptJson(ctx iris.Context) {
	p := Person{}
	if err := ctx.ReadJSON(&p); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
	}

	bytes, _ := json.Marshal(p)
	log.Print(string(bytes))
}
