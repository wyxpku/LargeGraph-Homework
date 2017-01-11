package controllers

import (
	"github.com/user/LargeGraph-Homework/models"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) Get() {
	c.TplName = "account.html"
}

func (c *AccountController) SignUp() {
	email := c.GetStringNotEmpty("email")
	name := c.GetStringNotEmpty("name")
	password := c.GetStringNotEmpty("password")
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if err := models.AddUser(email, name, password); err != nil {
		c.JsonError(err.Error())
		return
	} else if id, err := models.GetUserId(email, password); err != nil {
		c.JsonError(err.Error())
		return
	} else {
		c.SetSession("userId", id)
		// c.Redirect("/", 302)
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"id":      id,
		}
		c.ServeJSON()
	}
}

func (c *AccountController) Login() {
	email := c.GetStringNotEmpty("email")
	password := c.GetStringNotEmpty("password")
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if id, err := models.GetUserId(email, password); err != nil {
		c.JsonError(err.Error())
		return
	} else {
		c.SetSession("userId", id)
		// c.Redirect("/", 302)
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"id":      id,
		}
		c.ServeJSON()
	}
}

func (c *AccountController) Logout() {
	id := c.GetSession("userId")
	if id != nil {
		c.DelSession("userId")
	}
	c.Redirect("/account", 302)
}
