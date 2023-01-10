package models

import (
	"errors"
	"time"
	"strconv"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id           int64
	Username     string    `xorm:"varchar(200)" json:"username"`
	Nickname     string    `xorm:"varchar(200)" json:"nickname"`
	Phone     string    `xorm:"varchar(200)" json:"phone"`
	Email     string    `xorm:"varchar(200)" json:"email"`
	Salt         string    `xorm:"varchar(200)" json:"salt"`
	Age          int       `xorm:"int(2)" json:"age"`
	Avatar       string    `xorm:"TEXT" json:"avatar"`
	Loginfailure int       `xorm:"int(10)" json:"loginfailure"`
	Logintime    int       `xorm:"int(10)" json:"logintime"`
	Loginip      string    `xorm:"varchar(200)" json:"loginip"`
	Token        string    `xorm:"varchar(59)"`
	Password     string    `xorm:"varchar(200)"`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}
type Adminjson struct {
	Id           int64     `json:"id"`
	Gid           int64     `json:"gid"`
	Username     string    `json:"username"`
	Nickname     string    `json:"nickname"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar       string    `json:"avatar"`
	Groupname     string    `json:"groupname"`
	Created     time.Time    `json:"createtime"`
}
func (a *Admin) TableName() string {
	return "admin"
}
type AdminGroup struct {
    Admin `xorm:"extends"`
	Authaccess `xorm:"extends"`
    Authgroup `xorm:"extends"`
}

func (AdminGroup) TableName() string {
	return "admin"
}

// users := make([]UserGroup, 0)
// engine.Join("INNER", "group", "group.id = user.group_id").Find(&users)
//根据用户名密码查询用户
func SelectUserByUserName(userName string) (*Admin, error) {
	a := new(Admin)
	has, err := orm.Where("username = ?", userName).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}
//根据用户id找用户返回数据
func SelectAdminById(Id string) (*Admin, error) {
	a := new(Admin)
	id, _ := strconv.ParseInt(Id, 10, 64)
	has, err := orm.Where("id = ?", id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}
//add
func AddAdmin(a *Admin) error {
	_, err := orm.Insert(a)
	return err
}
//分页列表
//列表
func GetUserList(limit int,pagesize int,search string,order string) []*Adminjson {
   //fmt.Println("搜索关键词",search)
//    limit,_ := strconv.Atoi(limits)
//  
//   if pagesize-1<1 {
	  page:=pagesize-1
//   }
   var listadmin []*AdminGroup
   //拼接搜索分页查询语句
   var byorder string
   byorder= "a.id ASC"
   if order == "-id" {
		   byorder = "a.id DESC"
	}
   if search!=""{
	orm.Table("admin").Alias("a").
	Cols("a.*,ac.*, g.id,g.name").
	Join("INNER", []string{"auth_group_access", "ac"}, "ac.uid = a.id").
	Join("INNER", []string{"auth_group", "g"}, "g.id = ac.gid").
	Where("a.username like ?", "%"+search+"%").
	OrderBy(byorder).
	// Orderby(byorder).
	Limit(limit, limit*page).
	Find(&listadmin)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
   } else {
    orm.Table("admin").Alias("a").
	Cols("a.*,ac.*, g.id,g.name").
	Join("INNER", []string{"auth_group_access", "ac"}, "a.id = ac.uid").
	Join("INNER", []string{"auth_group", "g"}, "g.id = ac.gid").
	// Where("a.username like ?", "%"+search+"%").
	OrderBy(byorder).
	Limit(limit, limit*page).
	Find(&listadmin)
   }
//    fmt.Println(listadmin)
  	adminlist := []*Adminjson{}
	for _,v := range listadmin {
		// 结构体exends要用结构体里的结构体去读
		node := &Adminjson {
			Id:v.Admin.Id,
			Gid:v.Authgroup.Id,
			Username:v.Admin.Username,
			Nickname:v.Admin.Nickname,
			Avatar:v.Admin.Avatar,
			Phone:v.Admin.Phone,
			Email:v.Admin.Email,
			Groupname:v.Authgroup.Name,
			Created:v.Admin.Created,
		}
		adminlist= append(adminlist,node)
	}
	return adminlist
}
func GetUsertotal(search string) int64 {
	var num int64
	num=0;
	a:=new(Admin)
   if search!=""{
	 total,err := orm.Cols("id", "username").Where("username like ?", "%"+search+"%").Count(a)
	 if err==nil {
		 num=total

	 }
   } else {
     total,err := orm.Cols("id", "username").Count(a)
	 if err==nil {
		 num=total

	 }
   }
   return num
}
//删除
func DeleteAdmin(id int64) int {
	a := new(Admin)
	outnum, _ := orm.ID(id).Delete(a)
	return int(outnum)

}

//修改
func Upadmin(a *Admin) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}