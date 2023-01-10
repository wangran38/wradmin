package models

import (
	"errors"
	"time"
)

type News struct {
	Id      int64     `json:"id"`
	Categoryid     int       `json:"categroy_id"`
	Title    string    `json:"title" xorm:"varchar(200)"`
	Image   string    `json:"image" xorm:"TEXT "`
	Keywords   string  `json:"keywords" xorm:"TEXT "`
	Description   string  `json:"description" xorm:" TEXT "`
	Content   string  `json:"content" xorm:"LONGTEXT "`
	Isshow     int       `json:"isshow" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Weigh   int  `json:"weigh"`
	Status  int       `json:"status"`
}

func (a *News) TableName() string {
	return "news"
}

//根据用户id找用户返回数据
func SelectNewsid(Id int64) (*News, error) {
	a := new(News)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("组别菜单数据出错！")
	}
	return a, nil

}

//添加
func Addnews(a *News) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func UpNews(a *News) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
func GetNewsList(limit int, pagesize int, search *News, order string) []*News {
	var page int
	listdata := []*News{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}
	session := orm.Table("news")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	// fmt.Println(stringid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Categoryid > 0 {
		session = session.And("category_id", search.Categoryid)
	}

	var byorder string
	byorder = "id ASC"
	if order != "" {
		byorder = "id DESC"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func GetNewstotal(search *News) int64 {
	var num int64
	num = 0
	session := orm.Table("news")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Title != "" {
		name := "%" + search.Title + "%"
		session = session.And("title LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	if search.Isshow > 0 {
		session = session.And("isshow", search.Isshow)
		// session = session.And("pid", rules.Title)
	}
	a := new(News)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteNews(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(News)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectNewsByTitle(Title string) (*News, error) {
	a := new(News)
	has, err := orm.Where("title = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}