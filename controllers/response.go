package controllers

import (
	"GoPracticeItem/encryption"
	json2 "encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// Response 执行情况回复
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

//Ident 身份信息回复
type Ident struct {
	Code     int       `json:"code"`
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Creatat int64     `json:"creatat"`
	Role     string    `json:"role"`
	Message  string    `json:"message"`
}

//GetResponseJSON 转换回复信息为json格式
func GetResponseJSON(r Response)string{
	jsonbyte,err := json2.MarshalIndent(r, "", " ")
	if err != nil{
		fmt.Println("getjson wrong")
	}
	json := string(jsonbyte)
	return json
}

//GetIdentJSON 转换身份信息为Json格式
func GetIdentJSON(r Ident)string{
	jsonbyte,err := json2.MarshalIndent(r, "", " ")
	if err != nil{
		fmt.Println("getjson wrong")
	}
	json := string(jsonbyte)
	return json
}

//VerifyEmailFormat 使用正则表达式对邮箱判断
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//TokenVerify 对token进行认证
func TokenVerify(writer http.ResponseWriter, request *http.Request)error{
	request.ParseForm()
	head  := request.Header
	tokenStr := head.Get("Token")
	id,_  := strconv.Atoi(head.Get("Id"))

	if encryption.Getting(tokenStr,id) != nil{
		var R Response
		R.Code = 403
		R.Message = "token is wrong"
		fmt.Fprint(writer,GetResponseJSON(R))
		return encryption.Getting(tokenStr,id)
	}
	return nil
}


