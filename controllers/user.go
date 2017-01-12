package controllers

import (
	"fmt"
	"github.com/user/LargeGraph-Homework/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) FriendMomentJson() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	} else {
		userId = id.(int64)
	}
	bound := c.GetString("bound")

	c.Data["json"] = models.GetFriendMoment(userId, bound)
	c.ServeJSON()
}

func (c *UserController) Get() {
	var myId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	} else {
		myId = id.(int64)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	user, err := models.GetUser(userId)
	if err != nil {
		c.JsonError(err.Error())
		return
	}

	c.Data["user"] = user
	c.Data["moments"] = models.GetUserMoment(userId)
	c.Data["following"] = models.GetUserFollowing(userId)
	c.Data["followed"] = models.GetUserFollowed(userId)
	c.Data["isFollow"] = models.IsFollow(myId, userId)
	c.Data["common"] = models.CommonFriend(myId, userId)
	c.TplName = "user.html"
}

func (c *UserController) GetUserInfo() {
	var myId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	} else {
		myId = id.(int64)
    }
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	user, err := models.GetUser(userId)
	if err != nil {
		c.JsonError(err.Error())
		return
	}

	c.Data["json"] = map[string]interface{}{
		"user":    user,
		"success": true,
        "I_Follow_Him": models.IsFollow(myId, userId),
        "He_Follow_Me": models.IsFollow(userId, myId),
	}
	c.ServeJSON()
}

func (c *UserController) GetUserMoment() {
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}
	c.Data["json"] = models.GetUserMoment(userId)
	c.ServeJSON()
}

func (c *UserController) GetUserFollowing() {
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}
	c.Data["json"] = models.GetUserFollowing(userId)
	c.ServeJSON()
}

func (c *UserController) GetUserFollowed() {
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}
	c.Data["json"] = models.GetUserFollowed(userId)
	c.ServeJSON()
}

func (c *UserController) AddMoment() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
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
		c.Redirect("/login", 302)
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
		c.Redirect(fmt.Sprintf("/user?id=%d", targetId), 302)
	}
}

func (c *UserController) RemoveFollow() {
	var userId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
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
		c.Redirect(fmt.Sprintf("/user?id=%d", targetId), 302)
	}
}

func (c *UserController) GetCommonFriend() {
	var myId int64
	if id := c.GetSession("userId"); id == nil {
		c.Redirect("/login", 302)
	} else {
		myId = id.(int64)
	}
	userId := c.GetId()
	if c.paramErr != "" {
		c.JsonError(c.paramErr)
		return
	}

	c.Data["json"] = models.CommonFriend(myId, userId)
	c.ServeJSON()
}
