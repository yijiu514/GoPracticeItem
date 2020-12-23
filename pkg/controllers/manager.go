package controllers

import (
	"GoPracticeItem/pkg/models"
	"net/http"
)

//Manager 测试editor接口权限
func Manager(w http.ResponseWriter, r *http.Request) {

	//验证令牌
	err := TokenVerify(r)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	id := GetMessageID(r)

	//权限认证
	err = manager(id)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	w.WriteHeader(201)
}

//查询用户权限并判断
func manager(id int) error {

	r, err := models.QuerForEditor(id)
	if err != nil {
		return models.MysqlWrong
	}
	if r == "manager" {
		return models.Validation
	}
	return nil
}
