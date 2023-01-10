package models
import (
	"errors"
)
type Authaccess struct {
	Uid int64
	Gid int64 
}

func (a *Authaccess) TableName() string {
	return "auth_group_access"
}
//根据用户id找用户返回数据
func SelectAdminGid(Id int64) (*Authaccess, error) {
	a := new(Authaccess)
	has, err := orm.Where("uid = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到所在组别！")
	}
	return a, nil

}
func AddAuthaccess(a *Authaccess) error {
	_, err := orm.Insert(a)
	return err
}

//删除
func DeleteAuthaccess(uid int64) int {
	a := new(Authaccess)
	outnum, _ := orm.Where("uid = ?", uid).Delete(a)
	return int(outnum)

}

//修改
//修改
func UpAuthaccess(a *Authaccess) (int64,error)  {
	affected, err := orm.Where("uid = ?", a.Uid).Update(a)
	return affected, err

}