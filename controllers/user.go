package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

//User 分析url并跳转相应的功能函数
func User(writer http.ResponseWriter, request *http.Request)  {

	url := request.URL.Path
	fmt.Println(request.URL.Path)
	urlstring := strings.Split( url,"/")
	fmt.Println(urlstring)
	user := urlstring[1]
	fmt.Println(user)
	if user == "user"{
		id   := urlstring[2]
		fmt.Println(id)
		fun  := urlstring[3]
		fmt.Println(fun)
		switch{
		case fun == "lock":
			Lock(id,writer,request)
		case fun =="role":
			Role(id,writer,request)
		case fun == "password":
			PasswordSet(id,writer,request)
		}

 	}
}