package routers

import (
	"GoPracticeItem/pkg/controllers"
	"net/http"
)

func init() {

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/my/password", controllers.PasswordChange)
	http.HandleFunc("/my/identity", controllers.Identity)

	http.HandleFunc("/", controllers.User)

	http.HandleFunc("/test/editor", controllers.Editor)
	http.HandleFunc("/test/manager", controllers.Manager)
}
