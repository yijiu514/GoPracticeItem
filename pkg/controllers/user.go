package controllers

import (
	"net/http"
	"strings"
)

//User 分析url并跳转相应的功能函数
func User(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	urlstring := strings.Split(url, "/")
	user := urlstring[1]

	if user == "user" {
		id := urlstring[2]
		fun := urlstring[3]
		switch {
		case fun == "lock":
			Lock(id, w, r)
		case fun == "role":
			Role(id, w, r)
		case fun == "password":
			PasswordSet(id, w, r)
		}
	}
}
