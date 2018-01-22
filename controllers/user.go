package controllers

import(
	"time"
	"chain_admin.golang/models"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/logs"
)

type UserController struct{
	BaseController
}

func (self *UserController) UserList(){
	page, err := self.GetInt("p")
	if err != nil{
		page = 1
	}
	users, total := models.GetUserList(page, self.pageSize)
	self.Data["users"] = users
	self.Data["total"] = total
	if total < 1{
		self.Data["hasdata"] = false
	}else{
		self.Data["hasdata"] = true
	}

	paginator := pagination.SetPaginator(self.Ctx, self.pageSize, total)
	self.Data["paginator"] = paginator
	self.display()
}

func (self *UserController)UserAdd(){
	self.modal_display()
}

func (self *UserController)AjaxAdd(){
	user := new(models.User)
	user.Username = self.GetString("username")
	user.Email = self.GetString("email")
	user.Nickname = self.GetString("nickname")
	user.Realname = self.GetString("realname")
	user.Gender,_ = self.GetInt("gender")
	user.Phone = self.GetString("phone")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := models.UserAdd(user)
	if err != nil{
        self.ajaxMsg(err.Error(), MSG_ERR)
	}
    self.ajaxMsg("添加用户成功", MSG_OK)
}

func (self *UserController)UserStatusChange(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	status, err := self.GetInt("status")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	user, err := models.GetUserById(id) 
	logs.Info(user)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	user.Status = status
	err = user.Update("status")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("用户【"+user.Username+"】状态已修改成功", MSG_OK)
}