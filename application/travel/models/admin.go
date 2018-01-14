package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminServiceProvider struct{}

var AdminServer *AdminServiceProvider

type Admin struct {
	ID       int64  `orm:"column(id);pk"`
	Name     string `orm:"column(name)"	json:"name"`
	Password string `orm:"column(password)	json:"password""`
}

func (this *AdminServiceProvider) New(name string, password string) error {
	o := orm.NewOrm()
	sql := "INSERT INTO travel.admin (name,password) VALUES(?,?)"
	_, err := o.Raw(sql, name, password).Exec()
	return err
}

func (this *AdminServiceProvider) Login(name string, password string) error {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM travel.admin WHERE name=? AND password=?", name, password).Exec()
	return err
}

func (this *AdminServiceProvider) Logout() {

}
