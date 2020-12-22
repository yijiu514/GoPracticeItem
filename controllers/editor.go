package controllers

import (
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
)

//Editor 测试editor接口权限
func Editor (writer http.ResponseWriter, request *http.Request){

	//验证令牌
	if TokenVerify(writer,request) == nil{
		request.ParseForm()
		head := request.Header
		id,_  := strconv.Atoi(head.Get("id"))
		println("get id success  : ")
		println(id)
		json   :=  editor(id)
		fmt.Fprintf(writer, json)
	}
}

func editor(id int) string {
	var R Response
	//查询角色
	r,_:= models.QuerForEditor(id)
	if r  == "manager"{

		fmt.Println("Permission denied")
		R.Code = 403
		R.Message = "Permission denied"
		return  GetResponseJSON(R)
	}

	fmt.Println("Validation passed")
	R.Code = 201
	R.Message = "Validation passed"
	return  GetResponseJSON(R)
}
