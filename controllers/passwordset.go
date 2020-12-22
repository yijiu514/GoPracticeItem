package controllers

import (
	"GoPracticeItem/encryption"
	"GoPracticeItem/models"
	"fmt"
	"net/http"
)

//PasswordSet 密码重置
func PasswordSet (id string,writer http.ResponseWriter, request *http.Request){
	request.ParseForm()
	if TokenVerify(writer,request)!= nil{
		id    := request.Form.Get("id")
		println("get id success : "+id)

		json    := passwordSet(id)

		fmt.Fprintf(writer, json)
	}
}

func passwordSet(id string) string  {

	var R Response
	pwd,salt:= encryption.Md5Salt("123456",8)
	err := models.UpdateForPwdSet(id,pwd,salt)
 	if err != nil{
		fmt.Println("reset failed")
		R.Code = 401
		R.Message = "reset failed"
		return  GetResponseJSON(R)
	}

	R.Code = 205
	R.Message = "reset success"
	return  GetResponseJSON(R)

}