package controllers

import (
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
)


//Manager 测试manager接口权限
func Manager(writer http.ResponseWriter, request *http.Request){
	if TokenVerify(writer,request)!= nil{
		request.ParseForm()
		head := request.Header
		id,_  := strconv.Atoi(head.Get("id"))
		json     :=  manager(id)
		fmt.Fprintf(writer, json)
	}

}

//通过id查询角色并进行权限认证
func manager(id int) string {
	var R Response
	r,_ := models.QuerForEditor(id)

	if r  == "editor"{
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


