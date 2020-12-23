package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"net/http"
	"strconv"
	"time"
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
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	err = encryption.TokenIssue(email, w)
	if err != nil {
		models.ErrorJudge(w, err)
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
		return models.EmailFormat
	}

	//密码验证
	pwd, salt, lockat, err := models.QueryForLogin(email)
	if err != nil {
		return models.MysqlWrong
	}
	if encryption.Md5Stirng(password+salt) != pwd {
		return models.PassWordWrong
	}

	//锁定验证
	if lockat > time.Now().Unix() {
		return models.Locked
	}

	//登陆成功
	return nil
}

//推出功能
func quit(w http.ResponseWriter, r *http.Request) {

	id, token := getid(r)

	err := encryption.TokenVerify(id, token)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	err = deletesessionsalt(id)
	if err != nil {
		models.ErrorJudge(w, err)
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
	err = models.UpdateSessionsalt(id, salt)
	return err
}
