package models

import (
	"github.com/JonSnow47/beego/application/travel/utility"
	"github.com/astaxie/beego/orm"
	"time"
)

type AdminServiceProvider struct{}

var AdminServer *AdminServiceProvider

type Admin struct {
	ID       int64     `orm:"column(id);pk"`
	Name     string    `orm:"column(name)"	json:"name"`
	Password string    `orm:"column(password)	json:"password""`
	Created  time.Time `orm:"column(created)"`
}

func (this *AdminServiceProvider) New(name string, pass string) error {
	o := orm.NewOrm()

	hash, err := utility.GenerateHash(pass)
	if err != nil {
		return err
	}
	password := string(hash)
	sql := "INSERT INTO travel.admin (name,password,created) VALUES(?,?,?)"
	created := time.Now()
	_, err = o.Raw(sql, name, password, created).Exec()
	return err
}

func (this *AdminServiceProvider) Login(name string, password string) (bool, error) {
	o := orm.NewOrm()
	var pass string
	err := o.Raw("SELECT password FROM travel.admin WHERE name=?", name).QueryRow(&pass)
	if err != nil {
		return false, err
	}
	flag := utility.CompareHash([]byte(pass), password)
	return flag, err
}

func (this *AdminServiceProvider) AdminInfo(name string) error {
	o := orm.NewOrm()
	var admin Admin
	err := o.Raw("SELECT (id,name,created) FROM travel.admin WHERE name=?", name).QueryRow(admin)
	return err
}

/*
func (this *AdminServiceProvider) IsAdminExist(name string) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("SELECT name FROM travel.admin WHERE name = ?", name).Exec()
	num, _ := res.RowsAffected()
	return num, err
}
*/
