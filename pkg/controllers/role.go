package controllers

import (
	"GoPracticeItem/pkg/models"
	"net/http"
	"strconv"
)

//Role 修改用户身份
func Role(id string, w http.ResponseWriter, r *http.Request) {

	err := TokenVerify(r)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	userid, _ := strconv.Atoi(id)

	err = role(userid, getrole(r))
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

}

//修改角色信息写入数据库
func role(id int, role string) error {
	if role != "manager" && role != "editor" && role != "admin" {
		return models.RoleWrong
	}

	err := models.QueryID2(id)
	if err != nil {
		return err
	}
	err = models.UpdateForRole(id, role)
	if err != nil {
		return err
	}
	return nil
}

//获取role信息
func getrole(r *http.Request) (role string) {
	r.ParseForm()
	role = r.Form.Get("email")
	return
}
