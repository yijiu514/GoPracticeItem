package controllers

import (
	"GoPracticeItem/pkg/models"
	"errors"
	"log"
	"net/http"
	"strconv"
)

var (
	LockSelf = errors.New("admin can not lock self")
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
		w.WriteHeader(406)
		log.Println(err)
		return
	}

	err = models.UpdateLock(id, 99999999)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(205)
}

//判断id是否相同
func judge(id int, idself int) (err error) {
	if id == idself {
		return LockSelf
	}
	return nil
}

//解锁用户
func unlock(id int, w http.ResponseWriter, r *http.Request) {

	err := models.UpdateLock(id, 0)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
	}

	w.WriteHeader(205)
}
