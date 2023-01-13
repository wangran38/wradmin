package models

import (
	// "fmt"
	"errors"
	_ "strings"
	"time"
	// "reflect"
)

type Device struct {
	Id         int64     `json:"id"`
	Categroyid int64     `json:"categroyid"`                                                          //设备的分类id,留口
	Simid int64     `json:"simid"`                                                          //设备的分类id,留口
	Number      int64    `json:"number"`                                                                 //设备编号
	Name       string    `json:"name"`                                                                 //设备名称
	Image   string    `json:"image" xorm:"TEXT "` //设备的图片
	Remark     string    `json:"remark"`                                                               //设备的备注
	Factory   string   `json:"factory"`    //设备所属厂家
	Contactpeople   string   `json:"contact_people"`    //设备所属厂家联系人
	Phone   string   `json:"contact_phone"`    //设备所属厂家联系人电话
	Isopen     int       `json:"is_open" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"` //是否显示
	Created    time.Time `json:"createtime" xorm:"created int"`
	Updated    time.Time `json:"updatetime" xorm:"updated int"`
	Deletetime int       `json:"deletetime"`
	Weigh      int       `json:"weigh"`                     //排序
	Status     string    `json:"status" xorm:"varchar(40)"` //设备状态
}

func (a *Device) TableName() string {
	return "device"
}

//获取列表数据

func GetDeviceList(limit int, pagesize int, search *Device, order string) []*Device {
	// orm.DB()
	// listdata = *[]Authrule
	var page int
	listdata := []*Device{} //查询后输出的结构体数据
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}

	session := orm.Table("device")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Number > 0 {
		session = session.And("number", search.Number)
	}
	if search.Name != "" {
		title := "%" + search.Name + "%"
		session = session.And("name LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Factory != "" {
		factory := "%" + search.Factory + "%"
		session = session.And("factory LIKE ?", factory)
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

func GetDevicetotal(search *Device) int64 { //查询的总数量
	var num int64
	num = 0
	session := orm.Table("device")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Number > 0 {
		session = session.And("number", search.Number)
	}
	if search.Name != "" {
		title := "%" + search.Name + "%"
		session = session.And("name LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Factory != "" {
		factory := "%" + search.Factory + "%"
		session = session.And("factory LIKE ?", factory)
		// session = session.And("pid", rules.Title)
	}
	//这里可以加经纬度的公里范围查询条件，例如离最近的设备，算了，不管，
	a := new(Device)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func DeleteDevice(id int64) int { //删除传ID来
	a := new(Device)
	outnum, _ := orm.ID(id).Delete(a)
	return int(outnum)

}
func AddDevice(a *Device) error { //增加，反射到结构体表的列名插死他
	_, err := orm.Insert(a)
	return err
}
func UpDevice(a *Device) (int64, error) { //更新，传结构体过来，取结构体的ID然后更新（a.id）
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}

//判断设备名称是否重复的查找
func SelectDeviceByName(Name string) (*Device, error) {
	a := new(Device)
	has, err := orm.Where("name = ?", Name).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}
