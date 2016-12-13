package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/account", &controllers.AccountController{})
	beego.Router("/account/signup", &controllers.AccountController{}, "post:SignUp")
	beego.Router("/account/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/account/logout", &controllers.AccountController{}, "post:Logout")
}
