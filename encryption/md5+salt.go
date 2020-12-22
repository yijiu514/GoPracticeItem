package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//RandomString 生成随机字符串
func RandomString (length int ) string {

	randstr := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result,byte(r.Intn(len(randstr))))
	}

	return hex.EncodeToString(result)
}

//Md5Salt  对密码进行“md5+盐”加密，并返回加密后的密文和盐
func Md5Salt (pwd string, SaltNum int) (string, string){
	salt := RandomString(8)

	pwdstring := pwd + salt
	data := []byte(pwdstring)

	h := md5.New()
	h.Write(data)
	output := hex.EncodeToString(h.Sum(nil))

	return output,salt
}

// Md5Stirng 对密码进行md5加密并返回密文
func Md5Stirng (pwd string) (string){
	data := []byte(pwd)

	h := md5.New()
	h.Write(data)
	output := hex.EncodeToString(h.Sum(nil))

	return output
}