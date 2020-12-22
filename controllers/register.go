package controllers

import (
	"GoPracticeItem/encryption"
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Register 用户注册
func Register (writer http.ResponseWriter, request *http.Request){

	//获取注册信息
	request.ParseForm()
	email    := request.Form.Get("email")
	password := request.Form.Get("password")
	time := time.Now().Unix()

	fmt.Println("get email success : "+email)
	fmt.Println("get password sucess : "+password)
	fmt.Printf("get time  success : %d\n",time)

	//生成token的盐
	sessionsalt := encryption.RandomString(8)

	json   := register(email,password,time,sessionsalt)
	id := models.QueryID(email)
	token := encryption.TokenCreate(id,sessionsalt)
	writer.Header().Set("token",token)
	writer.Header().Set("id",strconv.Itoa(id))
	fmt.Fprintf(writer, json)
}


//注册信息写入数据库
func register(email string,password string,time int64,sessionsalt string)string{
	var R Response
	if VerifyEmailFormat(email) == false {
		R.Code = 401
		R.Message = "the email format is not standardized"
		return  GetResponseJSON(R)
	}

	pwd,salt := encryption.Md5Salt(password,8)

	err := models.InsertForRegister(email,pwd,salt,time,sessionsalt)

	if err != nil{
		R.Code = 409
		R.Message = "the email is exist"
		return  GetResponseJSON(R)
	}

	R.Code = 201
	R.Message = "register success"
	return  GetResponseJSON(R)
}