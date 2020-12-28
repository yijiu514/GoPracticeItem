package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"log"
	"net/http"
	"strconv"
)

//PasswordSet 密码重置
func PasswordSet(idstr string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if errors.Is(err, encryption.TokenWrong) && errors.Is(err, encryption.TokenEmpty) {
		w.WriteHeader(401)
		log.Println(err)
		return
	}

	pwd, salt := encryption.Md5Salt("123456", 8)

	id, _ := strconv.Atoi(idstr)
	err = models.UpdatePwdChange(id, pwd, salt)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(205)
}
