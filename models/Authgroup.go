package models

import (
	"errors"
	"time"
)

type Authgroup struct {
	Id      int64     `json:"id"`
	Pid     int       `json:"pid"`
	Name    string    `json:"name" xorm:"LONGTEXT "`
	Rules   string    `json:"rules" xorm:"LONGTEXT "`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Status  int       `json:"status"`
}
type Authgrouptree struct {
	Id      int64     `json:"id"`
	Pid     int       `json:"pid"`
	Name    string    `json:"name" xorm:"LONGTEXT "`
	Rules   string    `json:"rules" xorm:"LONGTEXT "`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Status  int       `json:"status"`
	Children  []*Authgrouptree
}
func (a *Authgroup) TableName() string {
	return "auth_group"
}
//获取树状数据
func Getgrouptree(pid int) []*Authgrouptree {
	m := new(Authgroup)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treegrouplist(pid)

}

//全部菜单
func (m *Authgroup) Treegrouplist(pid int) []*Authgrouptree {
	// menus := new(Authgroup)
	// // 	var a []Authrule
	var menus []*Authgroup
	orm.Where("pid = ?", pid).Find(&menus)
	treelist := []*Authgrouptree{}
	for _, v := range menus {
		child := v.Treegrouplist(int(v.Id))
		node := &Authgrouptree{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:  v.Name,
			Rules: v.Rules,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}
//根据用户id找用户返回数据
func SelectGidRule(Id int64) (*Authgroup, error) {
	a := new(Authgroup)
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
func Addgroup(a *Authgroup) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func Upgroup(a *Authgroup) (int64,error) {
	affected, err := orm.Id(a.Id).Cols("pid").Update(a)
	return affected, err

}
func GetgroupList(limit int, pagesize int, search string, order string) []*Authgroup {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Authgroup{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "id ASC"
	if order == "-id" {
		byorder = "id DESC"
	}
	orm.Table("auth_group").
		Where("name like ?", "%"+search+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

func Getgrouptotal(search string) int64 {
	var num int64
	num = 0
	a := new(Authgroup)
	total, err := orm.Cols("id", "name").Where("name like ?", "%"+search+"%").Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteGroup(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Authgroup)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectGroupByName(Name string) (*Authgroup, error) {
	a := new(Authgroup)
	has, err := orm.Where("name = ?", Name).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}