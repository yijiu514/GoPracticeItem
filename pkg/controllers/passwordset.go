package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"net/http"
)

//PasswordSet 密码重置
func PasswordSet(id string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	pwd, salt := encryption.Md5Salt("123456", 8)
	err = models.UpdateForPwdSet(id, pwd, salt)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}
	w.WriteHeader(205)
}
