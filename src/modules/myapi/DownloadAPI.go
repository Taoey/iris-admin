package myapi

import (
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
)

// 官方下载示例 https://www.studyiris.com/example/fileServer/sendFiles.html
// 推荐使用官方的代码进行下载，其实如果看源码的话，源码使用的下载模式和第二种相同
func ExcelDownloadDemo1(ctx iris.Context) {
	pwd, _ := os.Getwd()

	filedir := pwd + "/files/"
	filename := "data1.txt"
	filepath := filedir + filename

	ctx.SendFile(filepath, filename)
}

// 互联网示例：go实现上传和下载excel接口 https://blog.csdn.net/weixin_43456598/article/details/100696033

// 通过自己设置header的方式下载
func ExcelDownloadDemo2(ctx iris.Context) {
	pwd, _ := os.Getwd()

	filedir := pwd + "/files/"
	filename := "7z1900-x64.exe"
	filepath := filedir + filename

	f, _ := os.Open(filepath)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Write(data)
}

func ExcelDownloadDemo3(ctx iris.Context) {
	filename := "./data1.txt1"
	filepath := filename

	ctx.SendFile(filepath, "11.txt")
}
