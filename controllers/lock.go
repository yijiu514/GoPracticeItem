package controllers

import (
	"GoPracticeItem/models"
	"fmt"
    "net/http"
	"strconv"
)


//Lock 锁定或解锁用户
func Lock (id string,writer http.ResponseWriter, request *http.Request){

	//获取信息
	request.ParseForm()
	method := request.Method
	head   := request.Header
	idself,_ := strconv.Atoi(head.Get("id"))
	//POST请求则锁定用户
	if method == "POST"{
		if TokenVerify(writer,request)!= nil{
			time,_ := strconv.ParseInt("99999999",10,64)
			id, _ := strconv.Atoi(id)
			json := lock (id,idself,time)
			fmt.Fprintf(writer, json)
		}
	}
	//DELETE请求则解锁用户
	if method == "DELETE"{
		if TokenVerify(writer,request)!= nil{
			id, _ := strconv.Atoi(id)
			json := unlock(id)
			fmt.Fprintf(writer,json)
		}
	}
}
//锁定用户
func lock(id int,idself int,time int64)string{
	var R Response
	models.UpdateForLock(id,time)

	if id == idself{
		R.Code = 403
		R.Message = "can not lock yourself"
		return  GetResponseJSON(R)
	}

	fmt.Println("lock success")
	R.Code = 201
	R.Message = "lock success"
	return  GetResponseJSON(R)
}

//解锁用户
func unlock(id int)string{
	var R Response
	err := models.UpdateForLock(id,0)
	if err != nil{
		fmt.Println("unlock failed")
		R.Code = 401
		R.Message = "unlock failed"
		return  GetResponseJSON(R)
	}
	fmt.Println("unlock success")
	R.Code = 201
	R.Message = "unlock success"
	return  GetResponseJSON(R)
}