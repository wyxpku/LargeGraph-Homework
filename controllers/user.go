package controllers

import (
	"github.com/user/LargeGraph-Homework/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/account/", 302)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if user, err := models.GetUser(userId); err != nil {
		c.JsonError(err.Error())
		return
	} else {
		c.Data["user"] = user
	}
	if moments, err := models.GetUserMoment(userId); err != nil {
		c.JsonError(err.Error())
		return
	} else {
		c.Data["moments"] = moments
	}
	c.TplName = "user.html"
}

func (c *UserController) AddMoment() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/account", 302)
	} else {
		userId = id.(int64)
	}
	content := c.GetStringNotEmpty("content")
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if err := models.AddUserMoment(userId, content); err != nil {
		c.JsonError(err.Error())
	} else {
		c.Redirect("/", 302)
	}
}

func (c *UserController) AddFollow() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/account", 302)
	} else {
		userId = id.(int64)
	}
	targetId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if err := models.UserFollow(userId, targetId); err != nil {
		c.JsonError(err.Error())
	} else {
		c.JsonSuccess()
	}
}

func (c *UserController) RemoveFollow() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/account", 302)
	} else {
		userId = id.(int64)
	}
	targetId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	if err := models.UserUnfollow(userId, targetId); err != nil {
		c.JsonError(err.Error())
	} else {
		c.JsonSuccess()
	}
}
