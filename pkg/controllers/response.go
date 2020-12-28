package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var (
	//EmailWrong 邮箱格式错误
	EmailWrong = errors.New("the email does not conform to the format")
)

//VerifyEmailFormat 使用正则表达式对邮箱判断
func VerifyEmailFormat(email string) error {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(email) {
		return EmailWrong
	}
	return nil
}

//TokenVerify 对token进行认证
func TokenVerify(request *http.Request) error {
	request.ParseForm()
	head := request.Header
	tokenStr := head.Get("Token")
	id, _ := strconv.Atoi(head.Get("Id"))
	err := encryption.Getting(tokenStr, id)

	if err != nil {
		return fmt.Errorf("token verify failed %w", err)
	}

	return nil
}

//GetMessageID 获取id信息
func GetMessageID(r *http.Request) (id int) {
	r.ParseForm()
	head := r.Header
	id, _ = strconv.Atoi(head.Get("id"))
	return
}
