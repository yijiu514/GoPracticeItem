package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"log"
	"net/http"
	"strconv"
)

//PasswordChange 修改密码
func PasswordChange(w http.ResponseWriter, r *http.Request) {
	err := TokenVerify(r)
	if errors.Is(err, encryption.TokenWrong) && errors.Is(err, encryption.TokenEmpty) {
		w.WriteHeader(401)
		log.Println(err)
		return
	}

	id, newpassword := getidandnewpwd(r)
	pwd, salt := encryption.Md5Salt(newpassword, 8)
	err = models.UpdatePwdChange(id, pwd, salt)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(204)
}

func getidandnewpwd(r *http.Request) (id int, newpassword string) {
	r.ParseForm()
	head := r.Header
	id, _ = strconv.Atoi(head.Get("id"))
	newpassword = r.Form.Get("newpassword")
	return
}
