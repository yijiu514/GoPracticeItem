package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"log"
	"net/http"
	"strconv"
)

//PasswordSet 密码重置
func PasswordSet(idstr string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if err != nil {

		return
	}

	pwd, salt := encryption.Md5Salt("123456", 8)
	id, _ := strconv.Atoi(idstr)
	err = models.UpdatePwdChange(id, pwd, salt)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(205)
}
