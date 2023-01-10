package controllers

import (
	// "fmt"
	"wradmin/models"
	"time"
	"wradmin/utils"
	"github.com/gin-gonic/gin"
)

type Groupserch struct {
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
// type Any interface{}
//获取当前用户信息
func Getgrouplist(c *gin.Context) {
	//从header中获取到token
	var searchdata Groupserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	name := searchdata.Name
	order := searchdata.Order
	listdata := models.GetgroupList(limit, page, name, order)
	listnum := models.Getgrouptotal(name)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功1",
			"data":    result,
		})
		return
	}
}
//获取全部上下级
func TreeGroup(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token != "" || len(token) <= 0 {
		user := utils.GetLoginUser(token)
		access, _ := models.SelectAdminGid(user.Id)//查找组别
		if access == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "失败",
			"data": "",
		})
		return
		}
		// group, err := models.SelectGidRule(access.Gid)//查找组别菜单
		// if err != nil {
		// c.JSON(200, gin.H{
		// 	"code":    201,
		// 	"message": "失败",
		// 	"data": err,
		// })
		// return
		// }
	

        grouplist := models.Getgrouptree(0)


		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data": grouplist,
		})

	} else {
			c.JSON(200, gin.H{
			"code":    201,
			"message": "你没有权限，或已经退出",
			"data": "",
		})
	}
}
//删除用户组
func Delgroup(c *gin.Context) {
	var json models.Authgroup
	// json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	delnum := models.DeleteGroup(json.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "删除失败！",
			"data":    delnum,
		})
	}



}
//添加用户组
func AddGroup(c *gin.Context) {
	var formdata models.Authgroup
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Rulesdata := new(models.Authgroup)
	
	Rulesdata.Pid = formdata.Pid
	Rulesdata.Name = formdata.Name
	Rulesdata.Rules = formdata.Rules
	Rulesdata.Status = formdata.Status
	Rulesdata.Created = time.Now()
	info, _ := models.SelectGroupByName(Rulesdata.Name) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "组别已经存在！",
		})
		return
	}
	err := models.Addgroup(Rulesdata) //判断账号是否存在！
		if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		// result := make(map[string]interface{})
		// result["id"] = Rulestable.Id //返回当前总数
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据添加成功！",
			"data": "",
		})

	}
	
}

//修改用户组
func EditGroup(c *gin.Context) {
	var formdata models.Authgroup
	c.ShouldBind(&formdata)
	intodata := new(models.Authgroup)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Name = formdata.Name
	intodata.Rules = formdata.Rules
	if(formdata.Id<=0) {
	c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res,err := models.Upgroup(intodata) //判断账号是否存在！
		if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改数据出错！",
			"data": err,
		})
		return
	} else {
		
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据修改成功！",
			"data": res,
		})

	}
	}

}