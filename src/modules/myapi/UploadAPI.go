package myapi

import (
	"github.com/kataras/iris"
	. "github_com_Taoey_iris_cli/src/entity"
	"github_com_Taoey_iris_cli/src/modules/myservice"
	"io/ioutil"
)

func UploadAliBill(ctx iris.Context) {
	file, _, _ := ctx.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	s := string(bytes)
	myservice.UploadAliBillPrint(s)

	result := Message{
		Code: MESSAGE_OK,
	}
	ctx.JSON(result)
}
