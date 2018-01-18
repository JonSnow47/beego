package models

import (
	"fmt"
	"github.com/JonSnow47/beego/application/travel/utility"
	"github.com/astaxie/beego/orm"
	"time"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type UserServiceProvider struct{}

var UserServer *UserServiceProvider

type User struct {
	ID      int64     `orm:"column(id);pk"	json:"id"`
	Name    string    `orm:"column(name)"	json:"name"`
	Pass    string    `orm:"column(pass)"	json:"pass"`
	Created time.Time `orm:"column(created)"`
	Status  bool      `orm:"column(status)"`
}

func init() {
	orm.RegisterModel(new(User))
}

//Register
func (this *UserServiceProvider) Register(name string, pass string) error {
	o := orm.NewOrm()

	hash, err := utility.GenerateHash(pass)
	if err != nil {
		return err
	}
	password := string(hash)

	sql := "INSERT INTO travel.admin(name,password,created) VALUES(?,?,?)"
	created := time.Now()
	_, err = o.Raw(sql, name, password, created).Exec()
	return err
}

//Login
func (this *UserServiceProvider) Login(name string, password string) (bool,error) {
	o := orm.NewOrm()
	var pass string
	err := o.Raw("SELECT password FROM travel.admin WHERE name=?", name).QueryRow(&pass)
	if err != nil {
		return false, err
	}
	flag := utility.CompareHash([]byte(pass), password)
	return flag, err
}

//read user info
func (this *UserServiceProvider) ReadInfoByID(id int64) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.Raw("SELECT name,created FROM travel.user id=? status=1", id).QueryRow(user)
	return user, err
}

//delete account
func (this *UserServiceProvider) Delete(content User) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE travel.user SET status=0").Exec()
	return err
}
