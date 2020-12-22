package controllers

import (
	"GoPracticeItem/encryption"
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
)


//PasswordChange 修改密码
func  PasswordChange(writer http.ResponseWriter, request *http.Request){
	if TokenVerify(writer,request)!= nil{
		request.ParseForm()
		head  := request.Header
		id,_  := strconv.Atoi(head.Get("id"))
		newpassword := request.Form.Get("new_password")

		println("get password success : "+newpassword)

		code     := passwordChange(id,newpassword)
		fmt.Fprintf(writer, code)
	}

}

//更新密码到数据库
func passwordChange(id int,newpassword string)string{

	var R Response
	pwd,salt := encryption.Md5Salt(newpassword,8)
	err := models.UpdateForPwdChange(id,pwd,salt)
	if err != nil {
		fmt.Println("change password failed")
		R.Code = 401
		R.Message = "change password failed"
		return  GetResponseJSON(R)
	}

	fmt.Println("change password success")
	R.Code = 201
	R.Message = "change password success"
	return  GetResponseJSON(R)
}