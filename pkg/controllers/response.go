package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"net/http"
	"regexp"
	"strconv"
)

//VerifyEmailFormat 使用正则表达式对邮箱判断
func VerifyEmailFormat(email string) error {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(email) {
		return models.EmailFormat
	}
	return nil
}

//TokenVerify 对token进行认证
func TokenVerify(request *http.Request) error {
	request.ParseForm()
	head := request.Header
	tokenStr := head.Get("Token")
	id, _ := strconv.Atoi(head.Get("Id"))

	if encryption.Getting(tokenStr, id) != nil {
		return models.TokenWrong
	}
	return nil
}

//获取id信息
func GetMessageID(r *http.Request) (id int) {
	r.ParseForm()
	head := r.Header
	id, _ = strconv.Atoi(head.Get("id"))
	return
}
