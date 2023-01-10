package controllers

import (
	"wradmin/lib"
	"wradmin/models"
	"net/http"
	"time"
	// "fmt"
	// "strconv"
	"github.com/gin-gonic/gin"
)

type Adminform struct {
	Username string `form:"username" binding:"required" json:"username"`
	Nickname string `form:"nickname" binding:"required" json:"nickname"`
	Phone string `form:"phone" binding:"required" json:"phone"`
	Email string `form:"email" binding:"required" json:"email"`
	Avatar string `form:"avatar" binding:"required" json:"avatar"`
	Gid int64 `form:"gid" binding:"required" json:"gid"`
	Password string `form:"password" binding:"required" json:"password"`
}
type Adminserch struct {
	ID int `json:"id"`
	Username string `json:"title"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	Order string `json:"sort"`
}
// type AdminController struct {
// 	BaseController //继承统一判断是否登录或者是否有权限类
// 	// beego.Controller
// }
func AddAdmin(c *gin.Context) {
	var admindata Adminform
	if err := c.ShouldBind(&admindata); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"msg":     "表单未完整！",
			"msgcode": "-1",
			// "data":    reqIP,
		})
		return
	}
	// //读取数据库
	Admin := new(models.Admin)
	Admin.Username = admindata.Username
	Admin.Nickname = admindata.Nickname
	Admin.Phone = admindata.Phone
	Admin.Avatar = admindata.Avatar
	Admin.Email = admindata.Email
	info, _ := models.SelectUserByUserName(Admin.Username) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "该用户已经存在！",
		})
		return
	}
	pwd, salt := lib.Password(4, admindata.Password) //截取四位随机盐+上这个做原始密码
	Admin.Password = pwd
	Admin.Salt = salt
	Admin.Created = time.Now()
	err := models.AddAdmin(Admin) //判断账号是否存在！
	if err != nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		Authaccess := new(models.Authaccess)
		Authaccess.Uid = Admin.Id
		Authaccess.Gid = admindata.Gid
		err := models.AddAuthaccess(Authaccess) //判断账号是否存在！
		if err==nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "数据添加成功！",
			"data": "",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "数据添加失败！",
			"data": err,
		})

	}

	}
}
//修改
func EditAdmin(c *gin.Context) {
	var formdata models.Adminjson
	c.ShouldBind(&formdata)
	intodata := new(models.Admin)
	intodata.Id = formdata.Id
	intodata.Username = formdata.Username
	intodata.Nickname = formdata.Nickname
	intodata.Phone = formdata.Phone
	intodata.Email = formdata.Email
	intodata.Avatar = formdata.Avatar
	if(formdata.Id<=0) {
	c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res,err := models.Upadmin(intodata) //判断账号是否存在！
		if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改数据出错！",
			"data": err,
		})
		return
	} else {
		// var formdata1 models.Adminjson
		// c.ShouldBind(&formdata1)
		updata := new(models.Authaccess)
		updata.Uid = formdata.Id
		updata.Gid = formdata.Gid
		_,err := models.UpAuthaccess(updata) //判断账号是否存在！
		if err==nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "数据修改成功！",
				"data": res,
			})
		}


	}
	}

}
func GetAdminlist(c *gin.Context) {
	var searchdata Adminserch
	// c.BindJSON(&searchdata)
	c.ShouldBind(&searchdata)
	result := make(map[string]interface{})
	limit:= searchdata.Limit
	page:= searchdata.Page
	username:= searchdata.Username
	order:= searchdata.Order
		listdata := models.GetUserList(limit,page,username,order)
		listnum := models.GetUsertotal(username)
		
		result["page"] = page
		result["totalnum"] = listnum
		result["limit"] = limit
		if listdata== nil {
			c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data": "",
		})
		return
			} else {
result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功1",
			"data": result,
		})
		return
			}
}

func Deladmin(c *gin.Context) {
	var searchdata models.Adminjson
	c.BindJSON(&searchdata)
	delnum := models.DeleteAdmin(searchdata.Id)
	if delnum > 0 {
		del2 := models.DeleteAuthaccess(searchdata.Id)
		if del2 > 0 {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "删除成功！",
				"data":    del2,
			})
		} else {
			c.JSON(201, gin.H{
				"code":    201,
				"message": "删除用户关系失败！",
				"data":    del2,
			})
		}

	} else {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "删除失败！",
			"data":    delnum,
		})
	}

}