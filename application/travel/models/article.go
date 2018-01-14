package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ArticleServiceProvider struct{}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ID      int64     `orm:"column(id);pk"`
	Uid     int64     `orm:"column(u_id)"`
	Title   string    `orm:"column(title)""	json:"title"`
	Content string    `orm:"column(content)"	json:"content"`
	Created time.Time `orm:"column(created)"`
	status  bool      `orm:"column(status)"`
}

type View struct {
	Title   string
	Content string
}

func (this *ArticleServiceProvider) New(title string, content string) error {
	o := orm.NewOrm()
	sql := "INSERT INTO travel.article (title,content) VALUES(?,?)"
	_, err := o.Raw(sql, title, content).Exec()
	return err
}

func (this *ArticleServiceProvider) Read(id int64) (View, error) {
	o := orm.NewOrm()
	var article View
	sql := "SELECT title,content FROM travel.article WHERE id=?"
	err := o.Raw(sql, id).QueryRow(&article)
	return article, err
}

func (this *ArticleServiceProvider) Update() {

}

func (this *ArticleServiceProvider) Delete() {

}

func (this *ArticleServiceProvider) View(id int64) (View, error) {
	o := orm.NewOrm()
	var article View
	sql := "SELECT title,content FROM travel.article WHERE id=?"
	err := o.Raw(sql, id).QueryRow(&article)
	if err != nil && err != orm.ErrNoRows {
		o.Raw("UPDATE title.article SET views=views+1")
	}
	return article, err
}
