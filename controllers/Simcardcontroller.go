package controllers

import (
	"time"
	"wradmin/models"

	// "net/http"
	// "ginstudy/utils"
	"github.com/gin-gonic/gin"
)

type Simcardserch struct { //这个是查询结构体，你也可以用切片，我这里用的是固定的结构体反射
	Id       int64  `json:"id"`
	Iccid    int64  `json:"iccid"`    //与卡片上打印的ICCID一一对应（卡片上最后一位为英文字母，系统上最后一位会将这个英文字母随机转换成数字。另外系统只需根据卡片上的前19位数字进行搜索）
	Simid    int64  `json:"simid"`    //设备的分类id,SIM卡标识是物联网卡在系统的唯一标识，在API接口调用中需要用到
	Deviceid int64  `json:"deviceid"` //设备的id
	Msisdn   string `json:"msisdn"`   //相当于固定网的用户电话号码，是供用户拨打的公开号码。
	Imsi     string `json:"imsi"`     //国际移动用户识别码（IMSI），国际上为唯一识别一个移动用户所分配的号码。
	Name     string `json:"name"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Order    string `json:"order"`
}

// 获取当前用户信息
func GetSimcardlist(c *gin.Context) {
	//从header中获取到token，这里是获取从前端提交过来的参数绑定到结构体
	var searchdata Simcardserch
	c.BindJSON(&searchdata)
	// //读取数据库
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	result := make(map[string]interface{})
	// name:=""
	// fmt.Println(username)
	search := &models.Simcard{
		Id:       searchdata.Id,
		Iccid:    searchdata.Iccid,
		Simid:    searchdata.Simid,
		Deviceid: searchdata.Deviceid,
		Msisdn:   searchdata.Msisdn,
		Imsi:     searchdata.Imsi,
		Name:     searchdata.Name,
	}
	// fmt.Println(search.Title)
	listdata := models.GetSimcardList(limit, page, search, order)
	//这些

	listnum := models.GetSimcardtotal(search)

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

func DelSimcard(c *gin.Context) {
	var searchdata Simcardserch
	c.BindJSON(&searchdata)
	delnum := models.DeleteSimcard(searchdata.Id)
	if delnum > 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功！",
			"data":    delnum,
		})
	}

}
func AddSimcard(c *gin.Context) {
	var formdata models.Simcard
	c.ShouldBind(&formdata)

	Intodata := new(models.Simcard)

	Intodata.Iccid = formdata.Iccid
	Intodata.Simid = formdata.Simid
	Intodata.Deviceid = formdata.Deviceid
	Intodata.Name = formdata.Name
	Intodata.Image = formdata.Image
	Intodata.Remark = formdata.Remark
	Intodata.Msisdn = formdata.Msisdn
	Intodata.Imsi = formdata.Imsi
	Intodata.Isopen = formdata.Isopen

	Intodata.Created = time.Now()
	info, _ := models.SelectSimcardBySimid(Intodata.Simid) //判断SIMid是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "sim编号已存在！",
		})
		return
	}
	err := models.AddSimcard(Intodata) //判断账号是否存在！
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

func EditSimcard(c *gin.Context) {
	var formdata models.Simcard
	c.ShouldBind(&formdata)
	// 	c.JSON(200, gin.H{
	// 	"code": "201",
	// 	"msg":  "添加数据出错！",
	// 	"data": formdata,
	// })
	Editdata := new(models.Simcard)
	Editdata.Id = formdata.Id
	Editdata.Iccid = formdata.Iccid
	Editdata.Simid = formdata.Simid
	Editdata.Deviceid = formdata.Deviceid
	Editdata.Name = formdata.Name
	Editdata.Image = formdata.Image
	Editdata.Remark = formdata.Remark
	Editdata.Msisdn = formdata.Msisdn
	Editdata.Imsi = formdata.Imsi
	Editdata.Isopen = formdata.Isopen
	res, err := models.UpSimcard(Editdata) //判断账号是否存在！
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
