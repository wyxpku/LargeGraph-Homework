package main

import (
	"github.com/astaxie/beego"
	_ "github.com/user/LargeGraph-Homework/routers"
)

func main() {
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.Run()
}
