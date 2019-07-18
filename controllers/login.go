package controllers

import (
	"blog_api/models"
	"github.com/astaxie/beego"
)


type LoginController struct {
	beego.Controller
}

func (obj *LoginController) Login()  {
	var user = models.GetAllUsers()
	obj.Data["json"] = user
	obj.ServeJSON()
}

func (obj *LoginController) GetUserList() {
	var userList = models.GetAllUsers()
	obj.Data["json"] = userList
	obj.ServeJSON()
}