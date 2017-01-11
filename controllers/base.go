package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/models"
)

type BaseController struct {
	beego.Controller
	paramErr string
}

func (c *BaseController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (c *BaseController) GetStringNotEmpty(name string) string {
	if c.paramErr != "" {
		return ""
	}
	res := c.GetString(name)
	if res == "" {
		c.paramErr = fmt.Sprintf("This param should not be empty: %s", name)
	}
	return res
}

func (c *BaseController) GetId() int64 {
	if c.paramErr != "" {
		return -1
	}
	res, _ := c.GetInt64("id", -1)
	if res == -1 {
		c.paramErr = fmt.Sprintf("Invalid Id")
	}
	return res
}

func (c *BaseController) JsonSuccess() {
	c.Data["json"] = map[string]bool{"success": true}
	c.ServeJSON()
}

func (c *BaseController) JsonError(errStr string) {
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"error":   errStr,
	}
	c.ServeJSON()
}

func (c *BaseController) Get() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/account/", 302)
	} else {
		userId = id.(int64)
	}
	user, err := models.GetUser(userId)
	if err != nil {
		c.JsonError(err.Error())
		return
	}
	c.Data["user"] = user
	c.Data["following"] = models.GetUserFollowing(userId)
	c.Data["followed"] = models.GetUserFollowed(userId)
	c.Data["friendMoment"] = models.GetFriendMoment(userId, "")
	c.TplName = "index.html"
}

func (c *BaseController) GetUserJson() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.JsonError("No such user")
		return
	} else {
		userId = id.(int64)
	}
	user, err := models.GetUser(userId)
	if err != nil {
		c.JsonError(err.Error())
		return
	}
	c.Data["json"] = map[string]interface{}{
		"user":    user,
		"success": true,
	}
	c.ServeJSON()
}

func (c *BaseController) Admin() {
	users, err := models.GetUserAll()
	if err != nil {
		c.JsonError(err.Error())
		return
	}
	c.Data["users"] = users
	c.TplName = "admin.html"
}

func (c *BaseController) AdminJson() {
	users, err := models.GetUserAll()
	if err != nil {
		c.JsonError(err.Error())
		return
	}
	c.Data["json"] = users
	c.ServeJSON()
}
