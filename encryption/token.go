package encryption

import (
	"GoPracticeItem/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//Claims 创建claim结构体
type Claims struct {
	UserID int
	jwt.StandardClaims //设置claim信息结构体
}

//TokenCreate 生成token
func TokenCreate(id int,sessionsalt string) string {
	expirTime := time.Now().Add(2*time.Hour) //设置有效时间为2小时
	claims := &Claims{
		UserID: id,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "127.0.0.1",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString,err := token.SignedString([]byte(sessionsalt))

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(tokenString)
	return tokenString
}

//Getting 验证token
func Getting(tokenstring string,id int)error{
	if tokenstring == ""{
		fmt.Println("the token is nil")
		err := errors.New("the token is nil")
		return err
	}
	token,err := ParseToken(tokenstring,id)
	fmt.Println(tokenstring)
	if err != nil || !token.Valid{
		fmt.Println("the token is expired or wrong")
		return err
	}
	return nil
}

//ParseToken token解析
func ParseToken(tokenString string,id int) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) ( interface{}, error) {
		fmt.Println(models.QuerySessionSalt(id))
		return []byte(models.QuerySessionSalt(id)),nil
	})
	return token, err
}

// HeaderSet 将token设置到Header
func HeaderSet(TokenString string,w http.ResponseWriter){
	w.Header().Set("token", TokenString)
}

//HeaderIDSet 设置id到header
func HeaderIDSet(id string,w http.ResponseWriter){
	w.Header().Set("id", id)
}

//HeaderRemember 将密码和用户名和token设置到header
func HeaderRemember(email string,password string,token string,w http.ResponseWriter){
	headeremail := &http.Cookie{
		Name:   "email",
		Value:  email,
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	headerpassword := &http.Cookie{
		Name:   "password",
		Value:  password,
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	headertoken := &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	w.Header().Set("email", headeremail.String())
	w.Header().Set("password", headerpassword.String())
	w.Header().Set("token", headertoken.String())
}



