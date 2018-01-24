package controllers

import (
	"chain_admin.golang/util"
	"strconv"
	"github.com/astaxie/beego/logs"
	"strings"
	"github.com/astaxie/beego"

	"chain_admin.golang/models"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct{
	beego.Controller
	controllerName string
	actionName string
	pageSize int
	login_user *models.User
}

func (self *BaseController) Prepare(){
	self.pageSize = 10
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["appversion"] = beego.AppConfig.String("appversion")
	self.Data["appname"] = beego.AppConfig.String("appname")
	self.Data["cur_controller"] = self.controllerName
	self.Data["cur_action"] = self.actionName

	if self.auth(){
		self.Data["login_nickname"] = self.login_user.Nickname
	}
	
}

func (self *BaseController)auth() bool{
	auth := self.Ctx.GetCookie("auth")
	arr := strings.Split(auth, "|")
	if len(arr) == 2{
		idStr, authkey := arr[0], arr[1]
		id, _ := strconv.Atoi(idStr)
		user, _ := models.GetUserById(id)
		if authkey == util.Md5(self.getClientIp() + "|" + user.Password, false){
			self.login_user = user
			return true
		}
	}
	return false
}

func (self *BaseController) display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

func (self *BaseController) modal_display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/modal_layout.html"
	self.TplName = tplname
}

func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
	logs.Info("url: ",url)
	self.Redirect(url, 302)
	self.StopRun()
}
