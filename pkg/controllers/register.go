package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"net/http"
	"time"
)

//Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {

	//获取注册信息
	email, pwd, time := getmessage(r)

	//注册验证
	err := verifyregister(email, pwd)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	//写入数据库
	salt := encryption.RandomString(8)
	err = models.InsertForRegister(email, pwd, salt, time, salt)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	//下发令牌
	err = encryption.TokenIssue(email, w)
	if err != nil {
		models.ErrorJudge(w, err)
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
		return VerifyEmailFormat(email)
	}

	//数据库验证
	if models.QueryID(email) != nil {
		return models.QueryID(email)
	}

	return nil
}
