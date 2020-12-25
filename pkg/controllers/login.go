package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	PassWordWrong = errors.New("the password is wrong")
	Locked        = errors.New("the user is locked")
)

//Login 判断
func Login(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	switch {
	case method == "POST":
		login(w, r)
	case method == "DELETE":
		quit(w, r)
	}
}

//登陆功能
func login(w http.ResponseWriter, r *http.Request) {

	email, pwd := getform(r)

	err := verifylogin(email, pwd)
	if errors.Is(err, EmailWrong) {
		w.WriteHeader(401)
		return
	} else if errors.Is(err, PassWordWrong) {
		w.WriteHeader(201)
		return
	} else if errors.Is(err, Locked) {
		w.WriteHeader(201)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	err = encryption.TokenIssue(email, w)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(201)
}

//获取表格信息
func getform(r *http.Request) (email string, password string) {
	r.ParseForm()
	email = r.Form.Get("email")
	password = r.Form.Get("password")
	return
}

//进行登陆验证
func verifylogin(email string, password string) error {

	//邮箱验证
	if VerifyEmailFormat(email) != nil {
		return EmailWrong
	}

	//密码验证
	pwd, salt, lockat, err := models.QueryLogin(email)
	if err != nil {
		return fmt.Errorf("mysql wrong %w", err)
	}
	if encryption.Md5Stirng(password+salt) != pwd {
		return PassWordWrong
	}

	//锁定验证
	if lockat > time.Now().Unix() {
		return Locked
	}

	//登陆成功
	return nil
}

//退出功能
func quit(w http.ResponseWriter, r *http.Request) {

	id, token := getid(r)

	err := encryption.TokenVerify(id, token)
	if err != nil {
		//进行断言
		return
	}

	err = deletesessionsalt(id)
	if err != nil {
		//进行断言
		return
	}

}

//获取id和token
func getid(r *http.Request) (id int, token string) {
	head := r.Header
	id, _ = strconv.Atoi(head.Get("id"))
	token = head.Get("token")
	return id, token
}

//置session盐为0
func deletesessionsalt(id int) (err error) {
	salt := ""
	err = models.UpdateSessionSalt(id, salt)
	return fmt.Errorf("deletesessionsalt wrong %w", err)
}
