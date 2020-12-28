package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {

	//获取注册信息
	email, pwd, time := getmessage(r)

	//注册验证
	err := verifyregister(email, pwd)
	if errors.Is(err, EmailWrong) {
		w.WriteHeader(401)
		log.Println(err)
		return
	} else if errors.Is(err, models.UserNotExist) {
		w.WriteHeader(409)
		log.Println(err)
		return
	} else if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	//写入数据库
	salt := encryption.RandomString(8)
	err = models.InsertRegister(email, pwd, salt, time, salt)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	//下发令牌
	err = encryption.TokenIssue(email, w)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(201)
}

//信息获取
func getmessage(r *http.Request) (email string, password string, timenow int64) {
	r.ParseForm()
	email = r.Form.Get("email")
	password = r.Form.Get("password")
	timenow = time.Now().Unix()
	return
}

//注册验证
func verifyregister(email string, password string) (err error) {

	//邮箱验证
	if VerifyEmailFormat(email) != nil {
		return fmt.Errorf("register wrong %w", VerifyEmailFormat(email))
	}

	//数据库验证
	if models.IsUserExistN(email) != nil {
		return fmt.Errorf("register wrong %w", models.IsUserExistN(email))
	}

	return nil
}
