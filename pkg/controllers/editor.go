package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//UsrLMT 用户权限不足
var UsrLMT = errors.New("the user inadequate permissions")

//Editor 测试editor接口权限
func Editor(w http.ResponseWriter, r *http.Request) {

	//验证令牌
	err := TokenVerify(r)
	if errors.Is(err, encryption.TokenWrong) && errors.Is(err, encryption.TokenEmpty) {
		w.WriteHeader(401)
		log.Println(err)
		return
	}

	id := GetMessageID(r)

	//权限认证
	err = editor(id)
	if errors.Is(err, UsrLMT) {
		w.WriteHeader(403)
		return
	} else if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(201)
}

//查询用户权限并判断
func editor(id int) error {
	r, err := models.QueryRole(id)
	if err != nil {
		return fmt.Errorf("query role wrong %w", err)
	}
	if r == "manager" {
		return UsrLMT
	}
	return nil
}
