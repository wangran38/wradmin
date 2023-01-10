package utils

import (
	"encoding/json"
	"fmt"
	"wradmin/models"
	"log"
)

//json转map
func JsonToMap(jsonStr string) (jsonMap map[string]interface{}) {
	jsonMap = make(map[string]interface{}, 0)
	json.Unmarshal([]byte(jsonStr), &jsonMap)
	return
}

//通过redis用户的唯一标识获取用户信息
func GetLoginUserInfo(userName string) models.Admin {
	var user models.Admin
	jsonStr := GetValueByKey("loginAdmin_" + userName)
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		log.Println("json string to struct error", err)
	}
	return user
}

//通过查找数据库取得用户信息
func Getuserinfo(userName string) models.Admin {
	var user models.Admin
	Userinfo, err := models.SelectAdminById(userName)
	if err != nil {
		log.Println("无当前用户信息", err)
	}
	// json.Unmarshal([]byte(userinfo), &user)
	user = *Userinfo
	return user
}

//获取用户信息
func GetLoginUser(token string) models.Admin {
	var user models.Admin
	jsonStr := GetLoginUserName(token)
	userMap := JsonToMap(jsonStr)
	userName := userMap["userName"].(string)
	user = Getuserinfo(userName)
	return user
}

//字符串首字母转化大写 user -> User
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
