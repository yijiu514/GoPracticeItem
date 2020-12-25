package models

import (
	"errors"
	"fmt"
)

var (
	UserNotExist = errors.New("the user is not exist")
	UserISExist  = errors.New("the user is exist")
)

//QueryMaxID 找到数据库当前最大id
func QueryMaxID() (int, error) {
	var max int
	err := DB.QueryRow("select max(id) from user").Scan(&max)
	if err != nil {
		return -1, fmt.Errorf("Query wrong -> %w", err)
	}
	return max, nil
}

//IsUserExistY 查询用户是否存在
func IsUserExistY(email string) error {
	var num int
	err := DB.QueryRow("select count(*) from user where  email = ?", email).Scan(&num)
	if err != nil {
		return fmt.Errorf("Query wrong -> %w", err)
	}
	if num == 0 {
		return UserNotExist
	}
	return nil
}

//IsUserExistN 查询用户是否存在
func IsUserExistN(email string) error {
	var num int
	err := DB.QueryRow("select count(*) from user where  email = ?", email).Scan(&num)
	if err != nil {
		return fmt.Errorf("Query wrong -> %w", err)
	}
	if num == 1 {
		return UserISExist
	}
	return nil
}

//
func IDIsUserExistN(id int) error {
	var num int
	err := DB.QueryRow("select count (*)from user where  id = ?", id).Scan(&num)
	if err != nil {
		return fmt.Errorf("Query wrong -> %w", err)
	}
	if num == 0 {
		return UserNotExist
	}
	return nil
}

//QueryRole 根据id查询角色
func QueryRole(id int) (string, error) {
	var role string
	sqlStr := "select role from user where id = ?"
	err := DB.QueryRow(sqlStr, id).Scan(&role)
	if err != nil {
		return "", fmt.Errorf("Query wrong -> %w", err)
	}
	return role, nil
}

//QueryLogin 查询登陆需要的相关信息
func QueryLogin(str string) (string, string, int64, error) {
	var u User
	sqlStr := "select password,password_salt,lock_at from user where email = ?"
	err := DB.QueryRow(sqlStr, str).Scan(
		&u.password,
		&u.passwordsalt,
		&u.lockat,
	)
	if err != nil {
		return "", "", 0, fmt.Errorf("Query wrong -> %w", err)
	}
	return u.password, u.passwordsalt, u.lockat, nil
}

//QuerIdentity 查询身份相关信息
func QueryIdentity(id int) (string, int64, string, error) {
	var u User

	sqlStr := "select email,creat_at,role from user where id = ?"
	err := DB.QueryRow(sqlStr, id).Scan(
		&u.email,
		&u.creatat,
		&u.role,
	)
	if err != nil {
		return "", 0, "", fmt.Errorf("Query wrong -> %w", err)
	}
	return u.email, u.creatat, u.role, nil
}

//QueryIDandSessionSalt 根据email查salt
func QueryIDandSessionSalt(email string) (id int, salt string, err error) {
	sqlStr := "select id  sessionsalt from user where email = ?"
	err = DB.QueryRow(sqlStr, email).Scan(
		&id,
		&salt,
	)
	if err != nil {
		return 0, "", fmt.Errorf("Query wrong -> %w", err)
	}
	return id, salt, err
}

// QuerySessionSalt 根据id查询sessionsalt
func QuerySessionSalt(id int) (salt string, err error) {
	sqlStr := "select sessionsalt from user where id = ?"
	err = DB.QueryRow(sqlStr, id).Scan(&salt)
	if err != nil {
		return "", fmt.Errorf("Query wrong -> %w", err)
	}
	return salt, nil
}

//InsertRegister 插入注册信息
func InsertRegister(email string, password string, passwordsalt string, time int64, sessionsalt string) error {
	sqlStr := "insert into user(id,email,role,creat_at,lock_at,password,passwordalt,sessionsalt) value(?,?,'editor',?,0,?,?,?)"
	id, err := QueryMaxID()
	if err != nil {
		return fmt.Errorf("QueryMaxID wrong -> %w", err)
	}

	_, err = DB.Exec(sqlStr, id, email, time, password, passwordsalt, sessionsalt)
	if err != nil {
		return fmt.Errorf("Insert wrong -> %w", err)
	}
	return nil
}

//UpdateLock 更新锁定时间
func UpdateLock(id int, time int64) error {
	sqlStr := "update user set lock_at = ? where id = ?"
	_, err := DB.Exec(sqlStr, time, id)
	if err != nil {
		return fmt.Errorf("Update wrong -> %w", err)
	}
	return nil
}

//UpdatePwdChange 修改密码
func UpdatePwdChange(id int, password string, salt string) error {
	sqlStr := "update user set password = ?,password_salt = ?  where id = ?"
	_, err := DB.Exec(sqlStr, password, salt, id)
	if err != nil {
		return fmt.Errorf("Update wrong -> %w", err)
	}
	return nil
}

//UpdateRole 修改角色信息
func UpdateRole(id int, role string) error {
	sqlStr := "update user set role = ? where id = ?"
	_, err := DB.Exec(sqlStr, role, id)
	if err != nil {
		return fmt.Errorf("Update wrong -> %w", err)
	}
	return nil
}

//UpdateSessionSalt 更新sesscion_salt
func UpdateSessionSalt(id int, sessionsalt string) error {
	sqlStr := "update user set sessionsalt= ? where id = ?"
	_, err := DB.Exec(sqlStr, sessionsalt, id)
	if err != nil {
		return fmt.Errorf("Update wrong -> %w", err)
	}
	return nil
}
