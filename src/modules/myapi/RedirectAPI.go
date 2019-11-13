package myapi

import (
	"github.com/kataras/iris"
	"log"
)

func RedirectURL(ctx iris.Context) {
	log.Print(ctx.Path())
	ctx.Redirect("", 302)
}
