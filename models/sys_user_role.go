package models

import(
	"github.com/astaxie/beego/orm"
)

type UserRole struct{
	UserId int
	RoleId int `orm:"pk"`
}

func (self *UserRole) TableName() string{
	return "sys_user_role"
}

func UserRoleAdd(ur *UserRole)(int64, error){
    return orm.NewOrm().Insert(ur)
}

func UserRoleDeleteByUserId(userId int)(int64, error){
	return orm.NewOrm().QueryTable("sys_user_role").Filter("user_id", userId).Delete()
}