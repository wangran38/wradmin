package controllers

import (
	// "fmt"
	"wradmin/models"
	"time"
	// "wradmin/utils"
	"github.com/gin-gonic/gin"
)

type Categoryserch struct {
	Title  string `json:"title"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
// type Any interface{}
//获取当前用户信息
func Getcategorylist(c *gin.Context) {
	//从header中获取到token
	var searchdata Categoryserch
	c.BindJSON(&searchdata)
	// //读取数据库
	result := make(map[string]interface{})
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	title := searchdata.Title
	order := searchdata.Order
	listdata := models.GetcategoryList(limit, page, title, order)
	listnum := models.Getcategorytotal(title)

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
//获取全部上下级
func TreeCategory(c *gin.Context) {

        grouplist := models.Getcategorytree(0)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data": grouplist,
		})

}


// //添加用户组
func AddCategory(c *gin.Context) {
	var formdata models.Category
	c.ShouldBind(&formdata)
		// 	c.JSON(200, gin.H{
		// 	"code": "201",
		// 	"msg":  "添加数据出错！",
		// 	"data": formdata,
		// })
	Rulesdata := new(models.Category)
	
	Rulesdata.Pid = formdata.Pid
	Rulesdata.Title = formdata.Title
	Rulesdata.Image = formdata.Image
	Rulesdata.Keywords = formdata.Keywords
	Rulesdata.Description = formdata.Description
	Rulesdata.Content = formdata.Content
	Rulesdata.Created = time.Now()
	info, _ := models.SelectcategoryByTitle(Rulesdata.Title) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该分类已经存在！",
		})
		return
	}
	err := models.Addcategory(Rulesdata) //判断账号是否存在！
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
func EditCategory(c *gin.Context) {
	var formdata models.Category
	c.ShouldBind(&formdata)
	intodata := new(models.Category)
	intodata.Id = formdata.Id
	intodata.Pid = formdata.Pid
	intodata.Title = formdata.Title
	intodata.Image = formdata.Image
	intodata.Keywords = formdata.Keywords
	intodata.Description = formdata.Description
	intodata.Content = formdata.Content
	if(formdata.Id<=0) {
	c.JSON(201, gin.H{
			"code": 201,
			"msg":  "修改选择的ID出错！",
			"data": "",
		})
		return
	} else {
		res,err := models.Upcategory(intodata) //判断账号是否存在！
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

func DelCategory(c *gin.Context) {
	var searchdata models.Category
	c.BindJSON(&searchdata)
	delnum := models.Deletecategory(searchdata.Id)
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