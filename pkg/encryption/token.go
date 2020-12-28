package encryption

import (
	"GoPracticeItem/pkg/models"
	"errors"
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

var (
	//TokenEmpty 令牌为空
	TokenEmpty = errors.New("the token is empty")
	//令牌失效或者错误
	TokenWrong = errors.New("the token is wrong or expired")
)

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
		return "", fmt.Errorf("token encryption wrong %w", err)
	}
	return tokenstring, nil
}

//Getting 验证token
func Getting(tokenstring string, id int) error {
	if tokenstring == "" {
		return TokenEmpty
	}
	token, err := ParseToken(tokenstring, id)
	if err != nil || !token.Valid {
		return TokenWrong
	}
	return nil
}

//TokenVerify 令牌验证
func TokenVerify(id int, token string) (err error) {
	err = Getting(token, id)
	if err != nil {
		return fmt.Errorf("token verify failed %w", err)
	}
	return nil
}

//ParseToken token解析
func ParseToken(tokenString string, id int) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		salt, err := models.QuerySessionSalt(id)
		if err != nil {
			return []byte(""), fmt.Errorf("parse wrong %w", err)
		}
		return []byte(salt), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse wrong %w", err)
	}
	return token, nil
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
		return fmt.Errorf("query ID and sessionsalt wrong %w", err)
	}

	token, err := TokenCreate(id, salt)
	if err != nil {
		return fmt.Errorf("token create wrong %w", err)
	}

	idtoken := strconv.Itoa(id)
	HeaderSet(token, w)
	HeaderIDSet(idtoken, w)
	return nil
}
