package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ArticleServiceProvider struct{}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ID      int64     `orm:"column(id);pk"`
	Title   string    `orm:"column(title)"		json:"title"`
	Author  string    `orm:"column(author)"`
	Class   string    `orm:"column(class)"		json:"class""`
	Content string    `orm:"column(content)"	json:"content"`
	Created time.Time `orm:"column(created)"`
	Status  bool      `orm:"column(status)"`
}

type View struct {
	Title   string
	Author  string
	Class   string
	Content string
	Created string
}

func (this *ArticleServiceProvider) New(article Article, author string) error {
	o := orm.NewOrm()
	sql := "INSERT INTO travel.article (title,author,class,content,created) VALUES(?,?,?,?,?,?)"
	created := time.Now()
	_, err := o.Raw(sql, article.Title, author, article.Class, article.Content, created).Exec()
	return err
}

func (this *ArticleServiceProvider) Read(id int64) (View, error) {
	o := orm.NewOrm()
	var article View
	sql := "SELECT title,content FROM travel.article WHERE id=?"
	err := o.Raw(sql, id).QueryRow(&article)
	return article, err
}

func (this *ArticleServiceProvider) Update(article Article) error {
	o := orm.NewOrm()
	sql := "UPDATE travel.article SET title=? content=? WHERE status=1"
	_, err := o.Raw(sql, article.Title, article.Content).Exec()
	return err
}

func (this *ArticleServiceProvider) Delete() error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE travel.article status=0").Exec()
	return err
}

func (this *ArticleServiceProvider) View(id int64) (View, error) {
	o := orm.NewOrm()
	var article View
	sql := "SELECT title,content FROM travel.article WHERE id=?"
	err := o.Raw(sql, id).QueryRow(&article)
	if err != nil && err != orm.ErrNoRows {
		o.Raw("UPDATE travel.article SET views=views+1")
	}
	return article, err
}
