package controllers

import (
	"wradmin/models"
	_ "time"
    // "net/http"
    _ "strconv"
	// "wradmin/utils"
	"github.com/gin-gonic/gin"
)
type Cityserch struct {
	Id        int64  `json:"id"`
	Pid       int  `json:"pid"`
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"status" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    string    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Order     string `json:"sort"`
}
//获取当前用户信息
func Getcitylist(c *gin.Context) {
	//从header中获取到token
	var searchdata Cityserch
	c.BindJSON(&searchdata)
	// //读取数据库
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	result := make(map[string]interface{})
	search := &models.City{
		Id:        searchdata.Id,
		Pid:       searchdata.Pid,
		Shortname:     searchdata.Shortname,
		Level:  searchdata.Level,
		Pinyin: searchdata.Pinyin,
        Code: searchdata.Code,
        Zip: searchdata.Zip,
        First: searchdata.First,
	}
	// fmt.Println(search.Title)
	listdata := models.GetCityList(limit, page, search, order)
	tabledata := listdata
	// for _, v := range listdata {
	// 	node := &Rulestable{
	// 		Id:        v.Id,
	// 		Pid:       v.Pid,
	// 		Type:      v.Type,
	// 		Icon:      v.Icon,
	// 		Pathname:  v.Pathname,
	// 		Component: v.Component,
	// 		Title:     v.Title,
	// 		Ismenu:    v.Ismenu,
	// 		Weigh:     v.Weigh,
	// 		Status:    v.Status,
	// 		Created:   v.Created,
	// 	}
	// 	// node.Children = child
	// 	tabledata = append(tabledata, node)
	// }
	listnum := models.GetCitytotal(search)

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
