package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type UserServiceProvider struct{}

var UserServer *UserServiceProvider

type User struct {
	UserID   int64
	UserName string
}

func init() {
	orm.RegisterModel(new(User))
}

//Register
func (this *UserServiceProvider) Register(info *User) (int64, error) {
	o := orm.NewOrm()
	o.Using("travel")		//use databases:travel

	user := new(User)
	id, err := o.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, err
}

//read user info
func (this *UserServiceProvider) UserInfo() {
	return
}
