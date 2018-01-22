package models

import(
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct{
	Id int
	Username string
	Nickname string
	Realname string
	Gender int
	Email string
	Phone string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status int
}

func (self *User) TableName() string{
    return "sys_user"
}

func UserAdd(user *User)(int64,error){
	return orm.NewOrm().Insert(user)
}

func (user *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *User)UserDelete(fields ...string)(int64, error){
	return orm.NewOrm().Delete(user, fields...)
}

func GetUserById(id int)(*User, error){
	user := new(User)
	err := orm.NewOrm().QueryTable(user.TableName()).Filter("id", id).One(user)
    if err != nil{
		return nil, err
	}
	return user, nil
}

func GetUserList(page, pageSize int, filters ...interface{})([]*User, int64){
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable("sys_user")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (self User)GetStatus()string{
	if self.Status == 0{
		return "已停用"
	}else if self.Status == 1{
		return "已启用"
	}else{
		return "未知"
	}
}


