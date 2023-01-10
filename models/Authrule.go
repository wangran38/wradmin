package models

import (
	// "fmt"
	"strings"
	"time"
	"errors"
	// "reflect"
)

type Authrule struct {
	Id         int64     `json:"id"`
	Pid        int64     `json:"pid"`
	Type       string    `json:"type"`
	Icon       string    `json:"icon"`
	Pathname   string    `json:"pathname"`
	Title      string    `json:"title"`
	Remark     string    `json:"remark"`
	Ismenu     int       `json:"ismenu" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"`
	Created    time.Time `json:"createtime" xorm:"created int"`
	Updated    time.Time `json:"updatetime" xorm:"updated int"`
	Deletetime int       `json:"deletetime"`
	Weigh      int       `json:"weigh"`
	Status     string    `json:"status" xorm:"varchar(40)"`
	Component  string    `json:"component"`
}
type Treerule struct {
	Id        int64
	Pid       int64
	Type      string `json:"type"`
	Icon      string `json:"icon"`
	Pathname  string `json:"pathname"`
	Component string `json:"component"`
	Title     string `json:"title"`
	Remark    string
	Ismenu    int `json:"ismenu"`
	Weigh     int `json:"weigh"`
	Status    string `json:"status"`
	Children  []*Treerule
}

func (a *Authrule) TableName() string {
	return "auth_rule"
}

//获取树状数据
func Getruletree() []*Treerule {
	m := new(Authrule)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treelist(0)

}

//全部菜单
func (m *Authrule) Treelist(pid int64) []*Treerule {
	// menus := new(Authrule)
	// 	var a []Authrule
	var menus []Authrule
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Treerule{}
	for _, v := range menus {
		child := v.Treelist(v.Id)
		node := &Treerule{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Icon:      v.Icon,
			Pathname:  v.Pathname,
			Component: v.Component,
			Title:     v.Title,
			Remark:    v.Remark,
			Ismenu:    v.Ismenu,
			Weigh:     v.Weigh,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}

//获取用户树状数据
func Getruleadmintree(Rules string) []*Treerule {
	m := new(Authrule)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treelistgroup(0, Rules)

}

//传参得某个权限的菜单集合
func (m *Authrule) Treelistgroup(pid int64, Rules string) []*Treerule {
	// fmt.Println(Rules)
	//    a := new(Authrule)
	// // 	var a []Authrule
	var menus []*Authrule
	ids := strings.Split(Rules, ",") //转成数组用orm in
	// 	// ids:= string.Join(Rules,",")
	// Where("pid = ?", v.Id)
	orm.Where("pid = ?", pid).Where("status = ?", "normal").In("id", ids).Find(&menus)
	treelist := []*Treerule{}
	for _, v := range menus {
		child := v.Treelistgroup(v.Id, Rules)
		node := &Treerule{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Icon:      v.Icon,
			Pathname:  v.Pathname,
			Component: v.Component,
			Title:     v.Title,
			Remark:    v.Remark,
			Ismenu:    v.Ismenu,
			Weigh:     v.Weigh,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}

func GetRulesList(limit int, pagesize int, search *Authrule, order string) []*Authrule {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Authrule{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}

	session := orm.Table("auth_rule")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	// fmt.Println(stringid)
	// stringpid := strconv.FormatInt(search.Id, 10)
	// if search.Pid > 0 {
	// 	session = session.And("pid", search.Pid)
	// }
	// fmt.Println(stringpid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Component != "" {
		component := "%" + search.Component + "%"
		session = session.And("component LIKE ?", component)
		// session = session.And("pid", rules.Title)
	}
	if search.Pathname != "" {
		pathname := "%" + search.Pathname + "%"
		session = session.And("pathname LIKE ?", pathname)
		// session = session.And("pid", rules.Title)
	}
	var byorder string
	byorder = "id ASC"
	if order != "" {
		byorder = "id DESC"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata

}

func GetRulestotal(search *Authrule) int64 {
	var num int64
	num = 0
	session := orm.Table("auth_rule")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Title != "" {
		name := "%" + search.Title + "%"
		session = session.And("title LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	if search.Component != "" {
		name := "%" + search.Component + "%"
		session = session.And("component LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	if search.Pathname != "" {
		name := "%" + search.Pathname + "%"
		session = session.And("pathname LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	a := new(Authgroup)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteRules(id int64) int {
	a := new(Authrule)
	outnum, _ := orm.ID(id).Delete(a)
	return int(outnum)

}
func AddRules(a *Authrule) error {
	_, err := orm.Insert(a)
	return err
}
func UpRules(a *Authrule) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
//根据
func SelectRulesByTitle(Title string) (*Authrule, error) {
	a := new(Authrule)
	has, err := orm.Where("title = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}

// //根据用户名密码查询用户
// func SelectAllRules(Id int64) []Authrule {
// 	// a := new(Authrule)
// 	everyone := make([]Authrule, 0)
// 	orm.Where("Pid = ?", Id).Find(&everyone)
// 	for _,v:=range everyone {
// 				// var cc make([]rule,0)
// 				everyone.Children=orm.Where("Pid = ?", v.Id).Find(&everyone)
// 				// fmt.Println([]*rule.Children)
// 	}
// 	// if err != nil {
// 	// 	return nil, errors.New("没有数据！")
// 	// }

// 	return everyone

// }

//列表查询
// func RulesPageList(params *NoticeParam) ([]*Authrule, int64) {
// 	query := orm.NewOrm().QueryTable(NoticeTBName())
// 	data := make([]*Authrule, 0)
// 	//默认排序
// 	sortorder := "-id" //定义索引
// 	if len(params.Search.Value) > 0 {
// 		query = query.Filter("notice_id", params.Search.Value)
// 	}

// 	total, _ := query.Count()
// 	query.OrderBy(sortorder).Limit(params.Length, params.Start).All(&data)
// 	return data, total
// }
