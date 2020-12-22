package controllers

import (
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
)

//Role 修改用户身份
func Role (id string,writer http.ResponseWriter, request *http.Request){

	if TokenVerify(writer,request)!= nil{
		request.ParseForm()
		roles  := request.Form.Get("role")
		println("get role success : "+roles)
		id, _ := strconv.Atoi(id)
		json    := role(id,roles)
		fmt.Fprintf(writer, json)
	}

}

//修改角色信息写入数据库
func role(id int,role string) string {
	var R Response
	if role != "manager"&&role != "editor"&&role != "admin"{
		R.Code = 401
		R.Message = "the role is not exits"
		return GetResponseJSON(R)
	}

	err := models.UpdateForRole(id,role)
	if err != nil{
		R.Code = 404
		R.Message = "the user is not exits"
		return GetResponseJSON(R)
	}
	fmt.Println("update role success")
	R.Code  = 201
	R.Message = "update role success"
	return GetResponseJSON(R)
}