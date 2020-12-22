package controllers

import (
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
)

//Identity 查询用户信息并返回
func Identity (writer http.ResponseWriter, request *http.Request){

	var u Ident

	//令牌验证
	if TokenVerify(writer,request) == nil{
		//读取id
		request.ParseForm()
		head  := request.Header
		id,_  := strconv.Atoi(head.Get("id"))

		//信息写入结构体
		email,creat,role :=  models.QuerForIdentity(id)
		u.Code = 201
		u.ID = id
		u.Email = email
		u.Creatat =creat
		u.Role = role
		result :=  GetIdentJSON(u)
		fmt.Println(result)
		fmt.Fprintf(writer, result)
	}
}
