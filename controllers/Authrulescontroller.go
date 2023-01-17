package controllers

import (
	"wradmin/models"
	"time"
    // "net/http"
	// "wradmin/utils"
	"github.com/gin-gonic/gin"
)

type Rulesserch struct {
	Id        int64  `json:"id"`
	Pid       int64  `json:"pid"`
	Title     string `json:"title"`
	Pathname  string `json:"pathname"`
	Component string `json:"component"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Order     string `json:"order"`
}

type Rulestable struct {
	Id        int64     `json:"id"`
	Pid       int64     `json:"pid"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	Pathname  string    `json:"pathname"`
	Component string    `json:"component"`
	Type      string    `json:"type"`
	Ismenu    int       `json:"ismenu"`
	Weigh     int       `json:"weigh"`
	Status    string    `json:"status"`
	Created   time.Time `json:"createtime"`
}

//获取当前用户信息
func Getruleslist(c *gin.Context) {
	//从header中获取到token
	var searchdata Rulesserch
	c.BindJSON(&searchdata)
	// //读取数据库
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	result := make(map[string]interface{})
	// name:=""
	// fmt.Println(username)
	search := &models.Authrule{
		Id:        searchdata.Id,
		Pid:       searchdata.Pid,
		Title:     searchdata.Title,
		Pathname:  searchdata.Pathname,
		Component: searchdata.Component,
	}
	// fmt.Println(search.Title)
	listdata := models.GetRulesList(limit, page, search, order)
	tabledata := []*Rulestable{}
	for _, v := range listdata {
		node := &Rulestable{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Icon:      v.Icon,
			Pathname:  v.Pathname,
			Component: v.Component,
			Title:     v.Title,
			Ismenu:    v.Ismenu,
			Weigh:     v.Weigh,
			Status:    v.Status,
			Created:   v.Created,
		}
		// node.Children = child
		tabledata = append(tabledata, node)
	}
	listnum := models.GetRulestotal(search)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取列表数据失败",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = tabledata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}

func DelRules(c *gin.Context) {
	var searchdata Rulesserch
	c.BindJSON(&searchdata)
	delnum := models.DeleteRules(searchdata.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	}

}
func AddRules(c *gin.Context) {
	var formdata Rulestable
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Rulesdata := new(models.Authrule)
	
	Rulesdata.Pid = formdata.Pid
	Rulesdata.Title = formdata.Title
	Rulesdata.Pathname = formdata.Pathname
	Rulesdata.Icon = formdata.Icon
	Rulesdata.Component = formdata.Component
	Rulesdata.Created = time.Now()
	info, _ := models.SelectRulesByTitle(Rulesdata.Title) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "菜单名称已经存在！",
		})
		return
	}
	err := models.AddRules(Rulesdata) //判断账号是否存在！
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

func EditRules(c *gin.Context) {
	var formdata models.Authrule
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Rulesdata := new(models.Authrule)
	Rulesdata.Id = formdata.Id
	Rulesdata.Pid = formdata.Pid
	Rulesdata.Title = formdata.Title
	Rulesdata.Pathname = formdata.Pathname
	Rulesdata.Icon = formdata.Icon
	Rulesdata.Ismenu = formdata.Ismenu
	Rulesdata.Component = formdata.Component

	res,err := models.UpRules(Rulesdata) //判断账号是否存在！
		if err != nil {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改数据出错！",
			"data": err,
		})
		return
	} else {
		// result := make(map[string]interface{})
		// result["id"] = Rulestable.Id //返回当前总数
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新成功！",
			"data": res,
		})

	}
	
}
//获取当前用户 所有的权限控制
