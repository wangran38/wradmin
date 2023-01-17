package models

import (
	// "fmt"
	"errors"
	_ "strings"
	"time"
	// "reflect"
)

type Simcard struct {
	Id         int64     `json:"id"`
	Iccid int64     `json:"iccid"`                         //与卡片上打印的ICCID一一对应（卡片上最后一位为英文字母，系统上最后一位会将这个英文字母随机转换成数字。另外系统只需根据卡片上的前19位数字进行搜索）
	Simid int64     `json:"simid"`                         //设备的分类id,SIM卡标识是物联网卡在系统的唯一标识，在API接口调用中需要用到
	Deviceid int64     `json:"deviceid"`                         //设备的id
	Msisdn      string    `json:"msisdn"`                  //相当于固定网的用户电话号码，是供用户拨打的公开号码。
	Imsi       string    `json:"imsi"`                     //国际移动用户识别码（IMSI），国际上为唯一识别一个移动用户所分配的号码。
	Name       string    `json:"name"`                                                                 //SIM卡套餐名称
	Image   string    `json:"image" xorm:"TEXT "` //SIM卡的图片
	Remark     string    `json:"remark"`                                                               //SIM卡的备注
	Isopen     int       `json:"isopen" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"` //是否激活，
	Isbind     int       `json:"isbind" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"` //是否已绑定设备
	Weigh      int       `json:"weigh"`                     //排序
	Status     int       `json:"status" xorm:"not null default 1 comment('是否启用 默认1 未在线 2 在线 3 异常 ') TINYINT"` //是否显示
	Created    time.Time `json:"createtime" xorm:"created int"`
	Updated    time.Time `json:"updatetime" xorm:"updated int"`
	Deletetime int       `json:"deletetime"`
}

func (a *Simcard) TableName() string {
	return "Simcard"
}

//获取列表数据

func GetSimcardList(limit int, pagesize int, search *Simcard, order string) []*Simcard {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Simcard{} //查询后输出的结构体数据
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}

	session := orm.Table("Simcard")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Iccid >0 {
		session = session.And("iccid", search.Iccid)
	}
	if search.Deviceid >0 {
		session = session.And("deviceid", search.Deviceid)
	}
	if search.Name != "" {
		title := "%" + search.Name + "%"
		session = session.And("name LIKE ?", title)
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

func GetSimcardtotal(search *Simcard) int64 { //查询的总数量
	var num int64
	num = 0
	session := orm.Table("Simcard")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Iccid >0 {
		session = session.And("iccid", search.Iccid)
	}
	if search.Deviceid >0 {
		session = session.And("deviceid", search.Deviceid)
	}
	if search.Name != "" {
		title := "%" + search.Name + "%"
		session = session.And("name LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	//这里可以加经纬度的公里范围查询条件，例如离最近的设备，算了，不管，
	a := new(Simcard)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteSimcard(id int64) int { //删除传ID来
	a := new(Simcard)
	outnum, _ := orm.ID(id).Delete(a)
	return int(outnum)

}
func AddSimcard(a *Simcard) error { //增加，反射到结构体表的列名插死他
	_, err := orm.Insert(a)
	return err
}
func UpSimcard(a *Simcard) (int64, error) { //更新，传结构体过来，取结构体的ID然后更新（a.id）
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}

//判断是否唯一
func SelectSimcardBySimid(Simid int64) (*Simcard, error) {
	a := new(Simcard)
	has, err := orm.Where("simid = ?", Simid).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}
