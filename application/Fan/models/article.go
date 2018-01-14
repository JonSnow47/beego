package models

import (
	"github.com/astaxie/beego/orm"
)

type ArticleServiceProvider struct{}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ID      int64  `orm:"column(id);pk"`
	Title   string `orm:"column(title)"	json:"title"`
	Class   string `orm:"column(class)"	json:"class"`
	Content string `orm:"column(content)"	json:"content"`
	Views   int64  `orm:"column(views)"`
	Status  bool   `orm:"column(status)"`
	//DateTime orm.DateTimeField
}

type Display struct {
	Title   string `orm:"column(title)"`
	Class   string `orm:"column(class)"`
	Content string `orm:"column(content)"`
}

func Init() {
	orm.RegisterModel(new(Article))
}

//create
func (this *ArticleServiceProvider) New(title string, class string, content string) error {
	o := orm.NewOrm()
	sql := "INSERT INTO Fan.article (title,class,content) VALUES(?,?,?)"
	raw := o.Raw(sql, title, class, content)
	_, err := raw.Exec()
	return err
}

//update
func (this *ArticleServiceProvider) Update(article Article) error {
	o := orm.NewOrm()
	sql := "UPDATE article SET title=?, class=?, content=? WHERE id=? AND status=1;"
	values := []interface{}{article.Title, article.Class, article.Content, article.ID}
	_, err := o.Raw(sql, values).Exec()
	return err
}

//read
func (this *ArticleServiceProvider) Read(id int64) (Display, error) {
	o := orm.NewOrm()
	var display Display
	raw := o.Raw("SELECT title,class,content FROM Fan.article WHERE id=? AND status=1", id)
	err := raw.QueryRow(&display)
	return display, err
}

//delete
func (this *ArticleServiceProvider) Delete(id int64) error {
	o := orm.NewOrm()
	raw := o.Raw("UPDATE Fan.article SET status=0 WHERE id=? AND status=1", id)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) GetArticleID(title string) int64 {
	o := orm.NewOrm()
	var id int64
	o.Raw("SELECT id FROM article WHERE title=?", title).QueryRow(&id)
	return id
}
