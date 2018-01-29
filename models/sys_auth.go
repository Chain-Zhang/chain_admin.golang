package models

import(
	"github.com/astaxie/beego/orm"
	"time"
)

type Auth struct{
	Id int
	ParentId int
	Name string
	Desc string
	Url string
	Sort int
	Status int
	Icon string
	IsShow bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (self *Auth)TableName()string{
	return "sys_auth"
}

func AuthAdd(auth *Auth)(int64, error){
    return orm.NewOrm().Insert(auth)
}