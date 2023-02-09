package controllers

import (
	"time"
	"wradmin/models"

	// "net/http"
	// "ginstudy/utils"
	"github.com/gin-gonic/gin"
)

type Deviceserch struct { //这个是查询结构体，你也可以用切片，我这里用的是固定的结构体反射
	Id         int64  `json:"id"`
	Categroyid int64  `json:"categroyid"` //设备的分类id,留口
	Simid      int64  `json:"simid"`      //设备的分类id,留口
	Number     string `json:"number"`     //设备的分类id,留口
	Name       string `json:"name"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Order      string `json:"order"`
}

// 获取当前用户信息
func Getdevicelist(c *gin.Context) {
	//从header中获取到token，这里是获取从前端提交过来的参数绑定到结构体
	var searchdata Deviceserch
	c.BindJSON(&searchdata)
	// //读取数据库
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	result := make(map[string]interface{})
	// name:=""
	// fmt.Println(username)
	search := &models.Device{
		Id:         searchdata.Id,
		Categroyid: searchdata.Categroyid,
		Number:     searchdata.Number,
		Name:       searchdata.Name,
	}
	// fmt.Println(search.Title)
	listdata := models.GetDeviceList(limit, page, search, order)
	//这些是写路由的树形数据，要加进去就这样写

	listnum := models.GetDevicetotal(search)

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
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}

func DelDevice(c *gin.Context) {
	var searchdata Deviceserch
	c.BindJSON(&searchdata)
	delnum := models.DeleteDevice(searchdata.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	}

}
func AddDevice(c *gin.Context) {
	var formdata models.Device
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Intodata := new(models.Device)

	Intodata.Categroyid = formdata.Categroyid
	Intodata.Simid = formdata.Simid
	Intodata.Number = formdata.Number
	Intodata.Name = formdata.Name
	Intodata.Image = formdata.Image
	Intodata.Remark = formdata.Remark
	Intodata.Factory = formdata.Factory
	Intodata.Contactpeople = formdata.Contactpeople
	Intodata.Phone = formdata.Phone
	Intodata.Isopen = formdata.Isopen

	Intodata.Created = time.Now()
	info, _ := models.SelectDeviceByName(Intodata.Name) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "菜单名称已经存在！",
		})
		return
	}
	err := models.AddDevice(Intodata) //判断账号是否存在！
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

func EditDevice(c *gin.Context) {
	var formdata models.Device
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Editdata := new(models.Device)
	Editdata.Id = formdata.Id
	Editdata.Categroyid = formdata.Categroyid
	Editdata.Name = formdata.Name
	Editdata.Categroyid = formdata.Categroyid
	Editdata.Simid = formdata.Simid
	Editdata.Number = formdata.Number
	Editdata.Name = formdata.Name
	Editdata.Image = formdata.Image
	Editdata.Remark = formdata.Remark
	Editdata.Factory = formdata.Factory
	Editdata.Contactpeople = formdata.Contactpeople
	Editdata.Phone = formdata.Phone
	Editdata.Isopen = formdata.Isopen

	res, err := models.UpDevice(Editdata) //判断账号是否存在！
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
			"data": res,
		})

	}

}

//获取当前用户 所有的权限控制
