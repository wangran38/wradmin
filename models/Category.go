package models

import (
	"errors"
	"time"
)

type Category struct {
	Id      int64     `json:"id"`
	Pid     int       `json:"pid"`
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
type Categorytree struct {
	Id      int64     `json:"id"`
	Pid     int       `json:"pid"`
	Title    string    `json:"title" xorm:"LONGTEXT "`
	// Image   string    `json:"image" xorm:"LONGTEXT "`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Weigh   int  `json:"weigh"`
	Status  int       `json:"status"`
	Children  []*Categorytree
}
func (a *Category) TableName() string {
	return "category"
}
//获取树状数据
func Getcategorytree(pid int) []*Categorytree {
	m := new(Category)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treecategorylist(pid)

}

//全部菜单
func (m *Category) Treecategorylist(pid int) []*Categorytree {

	var menus []*Category
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Categorytree{}
	for _, v := range menus {
		child := v.Treecategorylist(int(v.Id))
		node := &Categorytree{
			Id:        v.Id,
			Pid:       v.Pid,
			Title:     v.Title,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}
//根据用户id找用户返回数据
func Selectcategoryid(Id int64) (*Category, error) {
	a := new(Category)
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
func Addcategory(a *Category) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func Upcategory(a *Category) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
func GetcategoryList(limit int, pagesize int, search string, order string) []*Category {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Category{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "id ASC"
	if order == "-id" {
		byorder = "id DESC"
	}
	orm.Table("category").
		Where("title like ?", "%"+search+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

func Getcategorytotal(search string) int64 {
	var num int64
	num = 0
	a := new(Category)
	total, err := orm.Cols("id", "title").Where("title like ?", "%"+search+"%").Count(a)
	if err == nil {
		num = total

	}
	return num
}

func Deletecategory(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Category)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectcategoryByTitle(Title string) (*Category, error) {
	a := new(Category)
	has, err := orm.Where("title = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}