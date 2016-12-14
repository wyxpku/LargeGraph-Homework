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
	c.TplName = "index.html"
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
