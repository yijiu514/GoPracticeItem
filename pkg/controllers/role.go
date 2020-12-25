package controllers

import (
	"GoPracticeItem/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	RoleWrong = errors.New("the role is not exist")
)

//Role 修改用户身份
func Role(id string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if err != nil {
		return
	}

	userid, _ := strconv.Atoi(id)

	err = role(userid, getrole(r))
	if errors.Is(err, RoleWrong) {
		w.WriteHeader(000)
		return
	}

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
