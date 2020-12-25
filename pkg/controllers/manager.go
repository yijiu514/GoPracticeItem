package controllers

import (
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//Manager 测试manager接口权限
func Manager(w http.ResponseWriter, r *http.Request) {

	//验证令牌
	err := TokenVerify(r)
	if err != nil {
		log.Println(err)
	}

	id := GetMessageID(r)

	//权限认证
	err = manager(id)
	if errors.Is(err, UsrLMT) {
		w.WriteHeader(423)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(201)
}

//查询用户权限并判断
func manager(id int) error {

	r, err := models.QueryRole(id)
	if err != nil {
		return fmt.Errorf("query role wrong %w", err)
	}
	if r == "editor" {
		return UsrLMT
	}
	return nil
}
