package controllers

import(
	"time"
	"chain_admin.golang/models"
)

type AuthController struct{
	BaseController
}

func (self *AuthController)List(){
	self.display()
}

func (self *AuthController) Add(){
	auth := new(models.Auth)
	auth.ParentId, _ = self.GetInt("parent_id")
	auth.Name = self.GetString("name")
	auth.Icon = self.GetString("icon")
	auth.Sort, _ = self.GetInt("sort")
	auth.Url = self.GetString("url")
	auth.Desc = self.GetString("desc")
	auth.Status = 1
	auth.IsShow = true
	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()

	_, err := models.AuthAdd(auth)
	if err != nil{
		self.ajaxMsg("添加权限失败", MSG_ERR)
	}
	self.ajaxMsg("添加权限成功", MSG_OK)
}

func (self *AuthController) GetList(){
	authList, _ := models.GetAuthList("status", 1)
    self.ajaxMsg(authList, MSG_OK)
}

func (self *AuthController) GetOne(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	auth, err := models.GetAuthById(id)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg(auth, MSG_OK)
}