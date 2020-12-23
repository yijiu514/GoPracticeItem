package controllers

import (
	"GoPracticeItem/pkg/models"
	json2 "encoding/json"
	"fmt"
	"net/http"
)

//Ident 身份信息
type Ident struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Creatat int64  `json:"creatat"`
	Role    string `json:"role"`
}

//Identity 查询用户信息并返回
func Identity(w http.ResponseWriter, r *http.Request) {

	//令牌验证
	err := TokenVerify(r)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	id := GetMessageID(r)
	messge, err := identitymessage(id)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}
	fmt.Fprintln(w, messge)

}

//GetIdentJSON 转换为Json输出
func GetIdentJSON(r Ident) string {
	jsonbyte, err := json2.MarshalIndent(r, "", " ")
	if err != nil {
		fmt.Println("getjson wrong")
	}
	json := string(jsonbyte)
	return json
}

//获取相关信息并返回结构体
func identitymessage(id int) (message string, err error) {
	var u Ident
	email, creat, role, err := models.QuerForIdentity(id)
	u.ID = id
	u.Email = email
	u.Creatat = creat
	u.Role = role
	return GetIdentJSON(u), err
}
