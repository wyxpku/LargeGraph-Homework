package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/user", &controllers.UserController{})

	beego.Router("/moment", &controllers.UserController{}, "POST:AddMoment")
}
