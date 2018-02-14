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
	IsMenu bool
	Method string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (self *Auth)TableName()string{
	return "sys_auth"
}


func AuthAdd(auth *Auth)(int64, error){
    return orm.NewOrm().Insert(auth)
}

func GetAuthList(filters ...interface{})([]*Auth, int64){
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable("sys_auth")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("sort").All(&list)
	return list, total
}

func GetAuthById(id int)(*Auth, error){
	auth := new(Auth)
	err := orm.NewOrm().QueryTable("sys_auth").Filter("id", id).One(auth)
    if err != nil{
		return nil, err
	}
	return auth, nil
}

func (auth *Auth) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(auth, fields...); err != nil {
		return err
	}
	return nil
}

func (auth *Auth)Delete(fields ...string)(int64, error){
	return orm.NewOrm().Delete(auth, fields...)
}