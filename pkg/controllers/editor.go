package controllers

import (
	"GoPracticeItem/pkg/models"
	"net/http"
)

//Editor 测试editor接口权限
func Editor(w http.ResponseWriter, r *http.Request) {

	//验证令牌
	err := TokenVerify(r)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	id := GetMessageID(r)

	//权限认证
	err = editor(id)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	w.WriteHeader(201)
}

//查询用户权限并判断
func editor(id int) error {

	r, err := models.QuerForEditor(id)
	if err != nil {
		return models.MysqlWrong
	}
	if r == "manager" {
		return models.Validation
	}
	return nil
}
