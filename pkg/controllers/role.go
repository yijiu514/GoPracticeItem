package controllers

import (
	"GoPracticeItem/pkg/encryption"
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	RoleWrong = errors.New("the role is not exist")
)

//Role 修改用户身份
func Role(id string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if errors.Is(err, encryption.TokenWrong) && errors.Is(err, encryption.TokenEmpty) {
		w.WriteHeader(401)
		log.Println(err)
		return
	}

	userid, _ := strconv.Atoi(id)

	err = role(userid, getrole(r))
	if errors.Is(err, RoleWrong) {
		w.WriteHeader(400)
		log.Println(err)
		return
	} else if errors.Is(err, models.UserNotExist) {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.WriteHeader(204)
}

//修改角色信息写入数据库
func role(id int, role string) error {
	if role != "manager" && role != "editor" && role != "admin" {
		return RoleWrong
	}

	err := models.IDIsUserExistN(id)
	if err != nil {
		return fmt.Errorf("idisuserexitsno wrong %w", err)
	}
	err = models.UpdateRole(id, role)
	if err != nil {
		return fmt.Errorf("updaterole wrong %w", err)
	}
	return nil
}

//获取role信息
func getrole(r *http.Request) (role string) {
	r.ParseForm()
	role = r.Form.Get("email")
	return
}
