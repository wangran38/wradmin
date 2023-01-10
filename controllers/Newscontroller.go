package controllers

import (
	// "fmt"
	"wradmin/models"
	"time"
	// "wradmin/utils"
	"github.com/gin-gonic/gin"
)
// type NewsController struct{}
type Newsserch struct {
	Id      int64     `json:"id"`
	Categoryid     int       `json:"categroy_id"`
	Title    string    `json:"title" xorm:"varchar(200)"`
	Image   string    `json:"image" xorm:"TEXT "`
	Keywords   string  `json:"keywords" xorm:"TEXT "`
	Description   string  `json:"description" xorm:" TEXT "`
	Content   string  `json:"content" xorm:"LONGTEXT "`
	Isshow     int       `json:"isshow" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Order     string `json:"order"`
}
// type Any interface{}
//获取当前用户信息
func Getnewslist(c *gin.Context) {
	//从header中获取到token
	var searchdata Newsserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	search := &models.News{
		Id:        searchdata.Id,
		Categoryid: searchdata.Categoryid,
		Title:     searchdata.Title,
		Isshow:  searchdata.Isshow,
	}
	listdata := models.GetNewsList(limit, page, search, order)
	listnum := models.GetNewstotal(search)

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
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}

// //添加用户组
func AddNews(c *gin.Context) {
	var formdata models.News
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Intodata := new(models.News)
	
	Intodata.Categoryid = formdata.Categoryid
	Intodata.Title = formdata.Title
	Intodata.Image = formdata.Image
	Intodata.Keywords = formdata.Keywords
	Intodata.Description = formdata.Description
	Intodata.Content = formdata.Content
	Intodata.Isshow = formdata.Isshow
	Intodata.Created = time.Now()
	info, _ := models.SelectNewsByTitle(Intodata.Title) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该分类已经存在！",
		})
		return
	}
	err := models.Addnews(Intodata) //判断账号是否存在！
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

// //修改用户组
func EditNews(c *gin.Context) {
	var formdata models.News
	c.ShouldBind(&formdata)
	updata := new(models.News)
	updata.Id = formdata.Id
	updata.Categoryid = formdata.Categoryid
	updata.Title = formdata.Title
	updata.Image = formdata.Image
	updata.Keywords = formdata.Keywords
	updata.Description = formdata.Description
	updata.Content = formdata.Content
	updata.Isshow = formdata.Isshow
	updata.Updated = time.Now()
	if(formdata.Id<=0) {
	c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res,err := models.UpNews(updata) //判断账号是否存在！
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

func DelNews(c *gin.Context) {
	var searchdata models.News
	c.BindJSON(&searchdata)
	delnum := models.DeleteNews(searchdata.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作失败！",
		})

	}

}