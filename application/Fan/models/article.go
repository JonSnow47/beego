package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ArticleServiceProvider struct{}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ID      int64  `orm:"column(id);pk"`
	Title   string `orm:"column(title)"	json:"title"`
	Author  string `orm:"column(author)"	json:"author"`
	Class   string `orm:"column(class)"	json:"class"`
	Content string `orm:"column(content)"	json:"content"`
	Views   int64  `orm:"column(views)"`
	Status  bool   `orm:"column(status)"`
	//DateTime orm.DateTimeField
}

type Display struct {
	Title   string `orm:"column(title)"`
	Author  string `orm:"column(author)"`
	Class   string `orm:"column(class)"`
	Content string `orm:"column(content)"`
}

func Init() {
	orm.RegisterModel(new(Article))
}
func (this *ArticleServiceProvider) New(title string, author string, class string, content string) error {
	o := orm.NewOrm()
	raw := o.Raw("INSERT INTO Fan.article (title,author,class,content) VALUES(?,?,?,?)", title, author, class, content)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) Update(article Display) error {
	o := orm.NewOrm()
	sql := "UPDATE Fan.article SET title=? AND class=? AND content=? WHERE status=1"
	values := []interface{}{article.Title, article.Class, article.Content}
	_, err := o.Raw(sql, values).Exec()
	return err
}

func (this *ArticleServiceProvider) Read(id int64) (Display, error) {
	o := orm.NewOrm()
	var display Display
	fmt.Println("dsd",id)
	raw := o.Raw("SELECT title,author,class,content FROM Fan.article WHERE id= ? AND status=1", id)

	err := raw.QueryRow(&display)
	return display, err
}

func (this *ArticleServiceProvider) Delete(title string, author string) error {
	o := orm.NewOrm()
	raw := o.Raw("UPDATE Fan.article SET status=0 WHERE title=? AND author=?", title, author)
	_, err := raw.Exec()
	return err
}
