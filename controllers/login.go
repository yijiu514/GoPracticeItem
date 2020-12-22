package controllers

import (
	"GoPracticeItem/encryption"
	"GoPracticeItem/models"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Login 实现登陆验证
func Login(writer http.ResponseWriter, request *http.Request){
	request.ParseForm()

	//获取请求类型
	method := request.Method
	println("get method success : "+method)

	if method =="POST"{

		//信息获取
		email     := request.Form.Get("email")
		password  := request.Form.Get("password")
		remember  := request.Form.Get("remember")

		url := request.URL
		fmt.Println(url)

		println("get email success : "+email)
		println("get password success : "+ password)
		println("get remember success : "+remember)


		//获取返回json信息
		json     := login(email,password)

		id    := models.QueryID(email)
		salt  := encryption.RandomString(8)
		token := encryption.TokenCreate(id,salt)

		//生成并下发令牌
		if remember != "1"{
			id := models.QueryID(email)
			fmt.Println(id)
			token :=encryption.TokenCreate(id,models.QuerySessionSalt(id))
			fmt.Println(models.QuerySessionSalt(id))
			encryption.HeaderSet(token,writer)
			encryption.HeaderIDSet(strconv.Itoa(id),writer)
			fmt.Fprintf(writer, json)
		}else {
			//如果记住密码，下发用户名密码
			fmt.Fprintf(writer, json)
			encryption.HeaderRemember(email,password,token,writer)
			encryption.HeaderIDSet(strconv.Itoa(id),writer)

		}

	//退出登陆则修改令牌盐
	if method =="DELETE"{
		err :=  models.UpdateSessionsalt(id,"")
		if err == nil{
			var R Response
			R.Code = 205
			R.Message = "quit success"
			fmt.Fprint(writer,GetResponseJSON(R))
		}
	}
	}
}

//登陆信息验证
func login(email string,password string) string {
	var R Response

	//邮箱验证
	if VerifyEmailFormat(email) == false {
		R.Code = 401
		R.Message = "the email format is not standardized"
		return GetResponseJSON(R)
	}

	pwd,salt,lockat := models.QueryForLogin(email)

	fmt.Println("get password from mysql success : "+pwd)
	fmt.Println("get password_salt from mysql success : "+salt)
	fmt.Println("user password : "+encryption.Md5Stirng(password+salt))

	//密码验证
	if (encryption.Md5Stirng(password+salt) != pwd){
		fmt.Println("the password is wrong ")
		R.Code = 401
		R.Message = "the password or email is wrong"
		return GetResponseJSON(R)
	}

	//判断用户是否被锁定
	if(lockat > time.Now().Unix()){
		fmt.Println("user is locked")
		R.Code = 423
		R.Message = "user is locked"
		return GetResponseJSON(R)
	}

	fmt.Println("login success")
	R.Code  = 201
	R.Message = "login success"
	return GetResponseJSON(R)
}