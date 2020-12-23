package controllers

import (
	"GoPracticeItem/pkg/models"
	"net/http"
	"strconv"
)

//Lock 锁定或解锁用户
func Lock(id string, w http.ResponseWriter, r *http.Request) {

	method := r.Method
	userid, _ := strconv.Atoi(id)
	switch {
	case method == "POST":
		lock(userid, w, r)
	case method == "DELETE":
		unlock(userid, w, r)
	}

}

//锁定用户
func lock(id int, w http.ResponseWriter, r *http.Request) {

	idself := GetMessageID(r)

	err := judge(id, idself)
	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	err = models.UpdateForLock(id, 99999999)

	if err != nil {
		models.ErrorJudge(w, err)
		return
	}
}

//判断id是否相同
func judge(id int, idself int) (err error) {
	if id == idself {
		return models.LockSelf
	}
	return nil
}

//解锁用户
func unlock(id int, w http.ResponseWriter, r *http.Request) {

	err := models.UpdateForLock(id, 0)

	if err != nil {
		models.ErrorJudge(w, err)
		return
	}

	w.WriteHeader(205)
}
