package models

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	EmailFormat      = errors.New("the email is wrong")
	TokenWrong       = errors.New("the token is wrong")
	TokenCreateWrong = errors.New("token creat wrong")
	Validation       = errors.New("Validation passed")
	MysqlWrong       = errors.New("mysql wrong")
	PassWordWrong    = errors.New("the password is wrong")
	Locked           = errors.New("the user is locked")
	LockSelf         = errors.New("can not lock you self")
	RoleWrong        = errors.New("the role is not exits ")
	UserExit         = errors.New("the user is exits")
)

//ErrorJudge 判断ERROR
func ErrorJudge(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, EmailFormat):
		w.WriteHeader(403)
	case errors.Is(err, TokenWrong):
		w.WriteHeader(401)
	case errors.Is(err, Validation):
		w.WriteHeader(403)
	case errors.Is(err, MysqlWrong):
		w.WriteHeader(213)
	case errors.Is(err, PassWordWrong):
		w.WriteHeader(401)
	case errors.Is(err, TokenWrong):
		w.WriteHeader(400)
	case errors.Is(err, Locked):
		w.WriteHeader(423)
	default:
		w.WriteHeader(456)
	}
}

//QueryMaxID 找到数据库当前最大id
func QueryMaxID() int {
	var max int
	err := DB.QueryRow("select max(id) from user").Scan(&max)
	if err != nil {
		fmt.Println("query failed")
		return 0
	}
	fmt.Println("query success")
	fmt.Printf("the max id is : %d\n", max)
	return max
}

//QueryID 查询是否存在
func QueryID(email string) error {
	var num int
	err := DB.QueryRow("select count (*)from user where  email = ?", email).Scan(&num)
	if err != nil {
		return MysqlWrong
	}
	if num > 0 {
		return UserExit
	}
	return nil
}

//QueryID2 根据id查询是否存在
func QueryID2(id int) error {
	var num int
	err := DB.QueryRow("select count (*)from user where  id = ?", id).Scan(&num)
	if err != nil {
		return MysqlWrong
	}
	if num > 0 {
		return UserExit
	}
	return nil
}

//QueryIDandSessionSalt 根据email查salt
func QueryIDandSessionSalt(email string) (id int, salt string, err error) {

	sqlStr := "select id  sessionsalt from user where email = ?"
	err = DB.QueryRow(sqlStr, email).Scan(
		&id,
		&salt,
	)
	if err != nil {
		return 0, "", err
	}
	return id, salt, err
}

//QuerForEditor 查询角色
func QuerForEditor(id int) (string, error) {
	var role string
	sqlStr := "select role from user where id = ?"
	err := DB.QueryRow(sqlStr, id).Scan(
		&role,
	)
	if err != nil {
		return "", err
	}
	return role, nil
}

//QueryForLogin 查询登陆需要的信息
func QueryForLogin(str string) (string, string, int64, error) {
	var u User
	sqlStr := "select password,password_salt,lock_at from user where email = ?"
	err := DB.QueryRow(sqlStr, str).Scan(
		&u.password,
		&u.passwordsalt,
		&u.lockat,
	)
	if err != nil {
		return "", "", 0, err
	}
	return u.password, u.passwordsalt, u.lockat, nil
}

//QuerForIdentity 查询身份相关信息
func QuerForIdentity(id int) (string, int64, string, error) {
	var u User

	sqlStr := "select email,creat_at,role from user where id = ?"
	err := DB.QueryRow(sqlStr, id).Scan(
		&u.email,
		&u.creatat,
		&u.role,
	)
	if err != nil {
		return "", 0, "", MysqlWrong
	}
	return u.email, u.creatat, u.role, nil
}

// QuerySessionSalt 根据id查询sessionsalt
func QuerySessionSalt(id int) string {
	var salt string
	sqlStr := "select sessionsalt from user where id = ?"
	err := DB.QueryRow(sqlStr, id).Scan(&salt)
	if err != nil {
		fmt.Println("query failed")
		fmt.Println(err)
		return "query failed"
	}
	fmt.Println("query success")
	return salt
}

//InsertForRegister 插入注册信息
func InsertForRegister(email string, password string, passwordsalt string, time int64, sessionsalt string) error {
	sqlStr := "insert into user(id,email,role,creat_at,lock_at,password,passwordalt,sessionsalt) value(?,?,'editor',?,0,?,?,?)"
	_, err := DB.Exec(sqlStr, QueryMaxID()+1, email, time, password, passwordsalt, sessionsalt)
	if err != nil {
		return MysqlWrong
	}
	return nil
}

//UpdateForLock 更新锁定时间
func UpdateForLock(id int, time int64) error {
	sqlStr := "update user set lock_at = ? where id = ?"
	_, err := DB.Exec(sqlStr, time, id)
	if err != nil {
		return MysqlWrong
	}
	return nil
}

//UpdateForPwdSet 充值密码
func UpdateForPwdSet(id string, password string, salt string) error {
	sqlStr := "update user set password = ?,password_salt = ?  where id = ?"
	_, err := DB.Exec(sqlStr, password, salt, id)
	if err != nil {
		return MysqlWrong
	}
	return nil
}

//UpdateForPwdChange 修改密码
func UpdateForPwdChange(id int, password string, salt string) error {
	sqlStr := "update user set password = ?,password_salt = ?  where id = ?"
	_, err := DB.Exec(sqlStr, password, salt, id)
	if err != nil {
		return MysqlWrong
	}
	return nil
}

//UpdateForRole 修改角色信息
func UpdateForRole(id int, role string) error {
	sqlStr := "update user set role = ? where id = ?"
	_, err := DB.Exec(sqlStr, role, id)
	if err != nil {
		return MysqlWrong
	}
	return nil
}

//UpdateSessionsalt 更新sesscion_salt
func UpdateSessionsalt(id int, sessionsalt string) error {
	sqlStr := "update user set sessionsalt= ? where id = ?"
	_, err := DB.Exec(sqlStr, sessionsalt, id)
	if err != nil {
		return MysqlWrong
	}
	return nil
}
