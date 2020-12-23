package encryption

import (
	"GoPracticeItem/pkg/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Claims 创建claim结构体
type Claims struct {
	UserID             int
	jwt.StandardClaims //设置claim信息结构体
}

//TokenCreate 生成token
func TokenCreate(id int, sessionsalt string) (tokenstring string, err error) {
	expirTime := time.Now().Add(2 * time.Hour) //设置有效时间为2小时
	claims := &Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err = token.SignedString([]byte(sessionsalt))

	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

//Getting 验证token
func Getting(tokenstring string, id int) error {
	if tokenstring == "" {
		return models.TokenWrong
	}
	token, err := ParseToken(tokenstring, id)
	if err != nil || !token.Valid {
		return models.TokenWrong
	}
	return nil
}

//TokenVerify 令牌验证
func TokenVerify(id int, token string) (err error) {

	if Getting(token, id) != nil {
		return Getting(token, id)
	}
	return nil
}

//ParseToken token解析
func ParseToken(tokenString string, id int) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(models.QuerySessionSalt(id))
		return []byte(models.QuerySessionSalt(id)), nil
	})
	return token, err
}

// HeaderSet 将token设置到Header
func HeaderSet(TokenString string, w http.ResponseWriter) {
	w.Header().Set("token", TokenString)
}

//HeaderIDSet 设置id到header
func HeaderIDSet(id string, w http.ResponseWriter) {
	w.Header().Set("id", id)
}

//TokenIssue 令牌下发
func TokenIssue(email string, w http.ResponseWriter) error {

	id, salt, err := models.QueryIDandSessionSalt(email)
	if err != nil {
		return models.MysqlWrong
	}
	token, err := TokenCreate(id, salt)
	if err != nil {
		return models.TokenCreateWrong
	}

	idtoken := strconv.Itoa(id)
	HeaderSet(token, w)
	HeaderIDSet(idtoken, w)
	return nil
}
